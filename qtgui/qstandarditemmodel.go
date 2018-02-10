package qtgui

// /usr/include/qt/QtGui/qstandarditemmodel.h
// #include <qstandarditemmodel.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 81
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin

type QStandardItemModel struct {
	*qtcore.QAbstractItemModel
}

func (this *QStandardItemModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemModel.GetCthis()
	}
}
func (this *QStandardItemModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemModel = qtcore.NewQAbstractItemModelFromPointer(cthis)
}
func NewQStandardItemModelFromPointer(cthis unsafe.Pointer) *QStandardItemModel {
	bcthis0 := qtcore.NewQAbstractItemModelFromPointer(cthis)
	return &QStandardItemModel{bcthis0}
}
func (*QStandardItemModel) NewFromPointer(cthis unsafe.Pointer) *QStandardItemModel {
	return NewQStandardItemModelFromPointer(cthis)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:326
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QStandardItemModel) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:330
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStandardItemModel(QObject *)
func NewQStandardItemModel(parent *qtcore.QObject /*777 QObject **/) *QStandardItemModel {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemModelFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:331
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStandardItemModel(int, int, QObject *)
func NewQStandardItemModel_1(rows int, columns int, parent *qtcore.QObject /*777 QObject **/) *QStandardItemModel {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModelC2EiiP7QObject", qtrt.FFI_TYPE_POINTER, rows, columns, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemModelFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:332
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStandardItemModel()
func DeleteQStandardItemModel(this *QStandardItemModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:336
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &)
func (this *QStandardItemModel) Index(row int, column int, parent *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:337
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex parent(const QModelIndex &)
func (this *QStandardItemModel) Parent(child *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg0 = child.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel6parentERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:339
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &)
func (this *QStandardItemModel) RowCount(parent *qtcore.QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:340
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &)
func (this *QStandardItemModel) ColumnCount(parent *qtcore.QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:341
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &)
func (this *QStandardItemModel) HasChildren(parent *qtcore.QModelIndex) bool {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:343
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &)
func (this *QStandardItemModel) Sibling(row int, column int, idx *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	var convArg2 = idx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:345
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int)
func (this *QStandardItemModel) Data(index *qtcore.QModelIndex, role int) *qtcore.QVariant /*123*/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:346
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)
func (this *QStandardItemModel) SetData(index *qtcore.QModelIndex, value *qtcore.QVariant, role int) bool {
	var convArg0 = index.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:348
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int)
func (this *QStandardItemModel) HeaderData(section int, orientation int, role int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:350
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setHeaderData(int, Qt::Orientation, const QVariant &, int)
func (this *QStandardItemModel) SetHeaderData(section int, orientation int, value *qtcore.QVariant, role int) bool {
	var convArg2 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:353
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)
func (this *QStandardItemModel) InsertRows(row int, count int, parent *qtcore.QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:354
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)
func (this *QStandardItemModel) InsertColumns(column int, count int, parent *qtcore.QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel13insertColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:355
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)
func (this *QStandardItemModel) RemoveRows(row int, count int, parent *qtcore.QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:356
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)
func (this *QStandardItemModel) RemoveColumns(column int, count int, parent *qtcore.QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel13removeColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:358
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &)
func (this *QStandardItemModel) Flags(index *qtcore.QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:359
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions()
func (this *QStandardItemModel) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:364
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QStandardItemModel) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:368
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)
func (this *QStandardItemModel) Sort(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:370
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * itemFromIndex(const QModelIndex &)
func (this *QStandardItemModel) ItemFromIndex(index *qtcore.QModelIndex) *QStandardItem /*777 QStandardItem **/ {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel13itemFromIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:371
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(const QStandardItem *)
func (this *QStandardItemModel) IndexFromItem(item *QStandardItem /*777 const QStandardItem **/) *qtcore.QModelIndex /*123*/ {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel13indexFromItemEPK13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:373
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * item(int, int)
func (this *QStandardItemModel) Item(row int, column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel4itemEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:374
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItem(int, int, QStandardItem *)
func (this *QStandardItemModel) SetItem(row int, column int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg2 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel7setItemEiiP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:375
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setItem(int, QStandardItem *)
func (this *QStandardItemModel) SetItem_1(row int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel7setItemEiP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:376
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * invisibleRootItem()
func (this *QStandardItemModel) InvisibleRootItem() *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel17invisibleRootItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:378
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * horizontalHeaderItem(int)
func (this *QStandardItemModel) HorizontalHeaderItem(column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel20horizontalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:379
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalHeaderItem(int, QStandardItem *)
func (this *QStandardItemModel) SetHorizontalHeaderItem(column int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel23setHorizontalHeaderItemEiP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:380
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * verticalHeaderItem(int)
func (this *QStandardItemModel) VerticalHeaderItem(row int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel18verticalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:381
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalHeaderItem(int, QStandardItem *)
func (this *QStandardItemModel) SetVerticalHeaderItem(row int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel21setVerticalHeaderItemEiP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:383
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalHeaderLabels(const QStringList &)
func (this *QStandardItemModel) SetHorizontalHeaderLabels(labels *qtcore.QStringList) {
	var convArg0 = labels.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel25setHorizontalHeaderLabelsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:384
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalHeaderLabels(const QStringList &)
func (this *QStandardItemModel) SetVerticalHeaderLabels(labels *qtcore.QStringList) {
	var convArg0 = labels.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel23setVerticalHeaderLabelsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:386
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowCount(int)
func (this *QStandardItemModel) SetRowCount(rows int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel11setRowCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:387
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnCount(int)
func (this *QStandardItemModel) SetColumnCount(columns int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel14setColumnCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:391
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void appendRow(QStandardItem *)
func (this *QStandardItemModel) AppendRow(item *QStandardItem /*777 QStandardItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel9appendRowEP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:395
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void insertRow(int, QStandardItem *)
func (this *QStandardItemModel) InsertRow(row int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel9insertRowEiP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:397
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool insertRow(int, const QModelIndex &)
func (this *QStandardItemModel) InsertRow_1(row int, parent *qtcore.QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel9insertRowEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:398
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool insertColumn(int, const QModelIndex &)
func (this *QStandardItemModel) InsertColumn(column int, parent *qtcore.QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel12insertColumnEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:400
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * takeItem(int, int)
func (this *QStandardItemModel) TakeItem(row int, column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel8takeItemEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:404
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * takeHorizontalHeaderItem(int)
func (this *QStandardItemModel) TakeHorizontalHeaderItem(column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel24takeHorizontalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:405
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * takeVerticalHeaderItem(int)
func (this *QStandardItemModel) TakeVerticalHeaderItem(row int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel22takeVerticalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:407
// index:0
// Public Visibility=Default Availability=Available
// [8] const QStandardItem * itemPrototype()
func (this *QStandardItemModel) ItemPrototype() *QStandardItem /*777 const QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel13itemPrototypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:408
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemPrototype(const QStandardItem *)
func (this *QStandardItemModel) SetItemPrototype(item *QStandardItem /*777 const QStandardItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel16setItemPrototypeEPK13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:414
// index:0
// Public Visibility=Default Availability=Available
// [4] int sortRole()
func (this *QStandardItemModel) SortRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel8sortRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:415
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortRole(int)
func (this *QStandardItemModel) SetSortRole(role int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel11setSortRoleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:417
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStringList mimeTypes()
func (this *QStandardItemModel) MimeTypes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel9mimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:419
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QStandardItemModel) DropMimeData(data *qtcore.QMimeData /*777 const QMimeData **/, action int, row int, column int, parent *qtcore.QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:422
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemChanged(QStandardItem *)
func (this *QStandardItemModel) ItemChanged(item *QStandardItem /*777 QStandardItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel11itemChangedEP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
