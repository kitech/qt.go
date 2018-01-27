package qtgui

// /usr/include/qt/QtGui/qstandarditemmodel.h
// #include <qstandarditemmodel.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 80
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QStandardItemModel) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:330
// index:0
// Public
// void QStandardItemModel(QObject *)
func NewQStandardItemModel(parent *qtcore.QObject /*777 QObject **/) *QStandardItemModel {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModelC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:331
// index:1
// Public
// void QStandardItemModel(int, int, QObject *)
func NewQStandardItemModel_1(rows int, columns int, parent *qtcore.QObject /*777 QObject **/) *QStandardItemModel {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModelC2EiiP7QObject", ffiqt.FFI_TYPE_VOID, cthis, rows, columns, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQStandardItemModelFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:332
// index:0
// Public virtual
// void ~QStandardItemModel()
func DeleteQStandardItemModel(*QStandardItemModel) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModelD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:336
// index:0
// Public virtual
// QModelIndex index(int, int, const QModelIndex &)
func (this *QStandardItemModel) Index(row int, column int, parent *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel5indexEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:337
// index:0
// Public virtual
// QModelIndex parent(const QModelIndex &)
func (this *QStandardItemModel) Parent(child *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = child.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel6parentERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:339
// index:0
// Public virtual
// int rowCount(const QModelIndex &)
func (this *QStandardItemModel) RowCount(parent *qtcore.QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel8rowCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:340
// index:0
// Public virtual
// int columnCount(const QModelIndex &)
func (this *QStandardItemModel) ColumnCount(parent *qtcore.QModelIndex) int {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel11columnCountERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:341
// index:0
// Public virtual
// bool hasChildren(const QModelIndex &)
func (this *QStandardItemModel) HasChildren(parent *qtcore.QModelIndex) bool {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel11hasChildrenERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:343
// index:0
// Public virtual
// QModelIndex sibling(int, int, const QModelIndex &)
func (this *QStandardItemModel) Sibling(row int, column int, idx *qtcore.QModelIndex) *qtcore.QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg2 = idx.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel7siblingEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:345
// index:0
// Public virtual
// QVariant data(const QModelIndex &, int)
func (this *QStandardItemModel) Data(index *qtcore.QModelIndex, role int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel4dataERK11QModelIndexi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:346
// index:0
// Public virtual
// bool setData(const QModelIndex &, const QVariant &, int)
func (this *QStandardItemModel) SetData(index *qtcore.QModelIndex, value *qtcore.QVariant, role int) bool {
	var convArg0 = index.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel7setDataERK11QModelIndexRK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:348
// index:0
// Public virtual
// QVariant headerData(int, Qt::Orientation, int)
func (this *QStandardItemModel) HeaderData(section int, orientation int, role int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel10headerDataEiN2Qt11OrientationEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), section, orientation, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:350
// index:0
// Public virtual
// bool setHeaderData(int, Qt::Orientation, const QVariant &, int)
func (this *QStandardItemModel) SetHeaderData(section int, orientation int, value *qtcore.QVariant, role int) bool {
	var convArg2 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:353
// index:0
// Public virtual
// bool insertRows(int, int, const QModelIndex &)
func (this *QStandardItemModel) InsertRows(row int, count int, parent *qtcore.QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel10insertRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:354
// index:0
// Public virtual
// bool insertColumns(int, int, const QModelIndex &)
func (this *QStandardItemModel) InsertColumns(column int, count int, parent *qtcore.QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel13insertColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:355
// index:0
// Public virtual
// bool removeRows(int, int, const QModelIndex &)
func (this *QStandardItemModel) RemoveRows(row int, count int, parent *qtcore.QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel10removeRowsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:356
// index:0
// Public virtual
// bool removeColumns(int, int, const QModelIndex &)
func (this *QStandardItemModel) RemoveColumns(column int, count int, parent *qtcore.QModelIndex) bool {
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel13removeColumnsEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:358
// index:0
// Public virtual
// Qt::ItemFlags flags(const QModelIndex &)
func (this *QStandardItemModel) Flags(index *qtcore.QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel5flagsERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:359
// index:0
// Public virtual
// Qt::DropActions supportedDropActions()
func (this *QStandardItemModel) SupportedDropActions() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel20supportedDropActionsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:364
// index:0
// Public
// void clear()
func (this *QStandardItemModel) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:368
// index:0
// Public virtual
// void sort(int, Qt::SortOrder)
func (this *QStandardItemModel) Sort(column int, order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel4sortEiN2Qt9SortOrderE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:370
// index:0
// Public
// QStandardItem * itemFromIndex(const QModelIndex &)
func (this *QStandardItemModel) ItemFromIndex(index *qtcore.QModelIndex) *QStandardItem /*777 QStandardItem **/ {
	var convArg0 = index.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel13itemFromIndexERK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:371
// index:0
// Public
// QModelIndex indexFromItem(const QStandardItem *)
func (this *QStandardItemModel) IndexFromItem(item *QStandardItem /*777 const QStandardItem **/) *qtcore.QModelIndex /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel13indexFromItemEPK13QStandardItem", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:373
// index:0
// Public
// QStandardItem * item(int, int)
func (this *QStandardItemModel) Item(row int, column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel4itemEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:374
// index:0
// Public
// void setItem(int, int, QStandardItem *)
func (this *QStandardItemModel) SetItem(row int, column int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg2 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel7setItemEiiP13QStandardItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:375
// index:1
// Public inline
// void setItem(int, QStandardItem *)
func (this *QStandardItemModel) SetItem_1(row int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel7setItemEiP13QStandardItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:376
// index:0
// Public
// QStandardItem * invisibleRootItem()
func (this *QStandardItemModel) InvisibleRootItem() *QStandardItem /*777 QStandardItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel17invisibleRootItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:378
// index:0
// Public
// QStandardItem * horizontalHeaderItem(int)
func (this *QStandardItemModel) HorizontalHeaderItem(column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel20horizontalHeaderItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:379
// index:0
// Public
// void setHorizontalHeaderItem(int, QStandardItem *)
func (this *QStandardItemModel) SetHorizontalHeaderItem(column int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel23setHorizontalHeaderItemEiP13QStandardItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:380
// index:0
// Public
// QStandardItem * verticalHeaderItem(int)
func (this *QStandardItemModel) VerticalHeaderItem(row int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel18verticalHeaderItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:381
// index:0
// Public
// void setVerticalHeaderItem(int, QStandardItem *)
func (this *QStandardItemModel) SetVerticalHeaderItem(row int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel21setVerticalHeaderItemEiP13QStandardItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:383
// index:0
// Public
// void setHorizontalHeaderLabels(const QStringList &)
func (this *QStandardItemModel) SetHorizontalHeaderLabels(labels *qtcore.QStringList) {
	var convArg0 = labels.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel25setHorizontalHeaderLabelsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:384
// index:0
// Public
// void setVerticalHeaderLabels(const QStringList &)
func (this *QStandardItemModel) SetVerticalHeaderLabels(labels *qtcore.QStringList) {
	var convArg0 = labels.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel23setVerticalHeaderLabelsERK11QStringList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:386
// index:0
// Public
// void setRowCount(int)
func (this *QStandardItemModel) SetRowCount(rows int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel11setRowCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), rows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:387
// index:0
// Public
// void setColumnCount(int)
func (this *QStandardItemModel) SetColumnCount(columns int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel14setColumnCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:391
// index:0
// Public inline
// void appendRow(QStandardItem *)
func (this *QStandardItemModel) AppendRow(item *QStandardItem /*777 QStandardItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel9appendRowEP13QStandardItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:395
// index:0
// Public inline
// void insertRow(int, QStandardItem *)
func (this *QStandardItemModel) InsertRow(row int, item *QStandardItem /*777 QStandardItem **/) {
	var convArg1 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel9insertRowEiP13QStandardItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:397
// index:1
// Public inline
// bool insertRow(int, const QModelIndex &)
func (this *QStandardItemModel) InsertRow_1(row int, parent *qtcore.QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel9insertRowEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:398
// index:0
// Public inline
// bool insertColumn(int, const QModelIndex &)
func (this *QStandardItemModel) InsertColumn(column int, parent *qtcore.QModelIndex) bool {
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel12insertColumnEiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:400
// index:0
// Public
// QStandardItem * takeItem(int, int)
func (this *QStandardItemModel) TakeItem(row int, column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel8takeItemEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:404
// index:0
// Public
// QStandardItem * takeHorizontalHeaderItem(int)
func (this *QStandardItemModel) TakeHorizontalHeaderItem(column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel24takeHorizontalHeaderItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), column)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:405
// index:0
// Public
// QStandardItem * takeVerticalHeaderItem(int)
func (this *QStandardItemModel) TakeVerticalHeaderItem(row int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel22takeVerticalHeaderItemEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), row)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:407
// index:0
// Public
// const QStandardItem * itemPrototype()
func (this *QStandardItemModel) ItemPrototype() *QStandardItem /*777 const QStandardItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel13itemPrototypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:408
// index:0
// Public
// void setItemPrototype(const QStandardItem *)
func (this *QStandardItemModel) SetItemPrototype(item *QStandardItem /*777 const QStandardItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel16setItemPrototypeEPK13QStandardItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:414
// index:0
// Public
// int sortRole()
func (this *QStandardItemModel) SortRole() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QStandardItemModel8sortRoleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:415
// index:0
// Public
// void setSortRole(int)
func (this *QStandardItemModel) SetSortRole(role int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel11setSortRoleEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), role)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:419
// index:0
// Public virtual
// bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QStandardItemModel) DropMimeData(data *qtcore.QMimeData /*777 const QMimeData **/, action int, row int, column int, parent *qtcore.QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:422
// index:0
// Public
// void itemChanged(QStandardItem *)
func (this *QStandardItemModel) ItemChanged(item *QStandardItem /*777 QStandardItem **/) {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QStandardItemModel11itemChangedEP13QStandardItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
