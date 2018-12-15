// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qdirmodel.h
// #include <qdirmodel.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

/*

 */
type QDirModel struct {
	*qtcore.QAbstractItemModel
}
type QDirModel_ITF interface {
	qtcore.QAbstractItemModel_ITF
	QDirModel_PTR() *QDirModel
}

func (ptr *QDirModel) QDirModel_PTR() *QDirModel { return ptr }

func (this *QDirModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemModel.GetCthis()
	}
}
func (this *QDirModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemModel = qtcore.NewQAbstractItemModelFromPointer(cthis)
}
func NewQDirModelFromPointer(cthis unsafe.Pointer) *QDirModel {
	bcthis0 := qtcore.NewQAbstractItemModelFromPointer(cthis)
	return &QDirModel{bcthis0}
}
func (*QDirModel) NewFromPointer(cthis unsafe.Pointer) *QDirModel {
	return NewQDirModelFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QDirModel) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdirmodel.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDirModel(const QStringList &, QDir::Filters, QDir::SortFlags, QObject *)

/*
Constructs a new directory model with the given parent. Only those files matching the nameFilters and the filters are included in the model. The sort order is given by the sort flags.
*/
func (*QDirModel) NewForInherit(nameFilters qtcore.QStringList_ITF, filters int, sort int, parent qtcore.QObject_ITF /*777 QObject **/) *QDirModel {
	return NewQDirModel(nameFilters, filters, sort, parent)
}
func NewQDirModel(nameFilters qtcore.QStringList_ITF, filters int, sort int, parent qtcore.QObject_ITF /*777 QObject **/) *QDirModel {
	var convArg0 unsafe.Pointer
	if nameFilters != nil && nameFilters.QStringList_PTR() != nil {
		convArg0 = nameFilters.QStringList_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg3 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModelC2ERK11QStringList6QFlagsIN4QDir6FilterEES3_INS4_8SortFlagEEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, filters, sort, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDirModel")
	return gothis
}

// /usr/include/qt/QtWidgets/qdirmodel.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDirModel(const QStringList &, QDir::Filters, QDir::SortFlags, QObject *)

/*
Constructs a new directory model with the given parent. Only those files matching the nameFilters and the filters are included in the model. The sort order is given by the sort flags.
*/
func (*QDirModel) NewForInheritp(nameFilters qtcore.QStringList_ITF, filters int, sort int) *QDirModel {
	return NewQDirModelp(nameFilters, filters, sort)
}
func NewQDirModelp(nameFilters qtcore.QStringList_ITF, filters int, sort int) *QDirModel {
	var convArg0 unsafe.Pointer
	if nameFilters != nil && nameFilters.QStringList_PTR() != nil {
		convArg0 = nameFilters.QStringList_PTR().GetCthis()
	}
	// arg: 3, QObject *=Pointer, QObject=Record, , Invalid
	var convArg3 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModelC2ERK11QStringList6QFlagsIN4QDir6FilterEES3_INS4_8SortFlagEEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, filters, sort, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDirModel")
	return gothis
}

// /usr/include/qt/QtWidgets/qdirmodel.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDirModel(QObject *)

/*
Constructs a new directory model with the given parent. Only those files matching the nameFilters and the filters are included in the model. The sort order is given by the sort flags.
*/
func (*QDirModel) NewForInherit1(parent qtcore.QObject_ITF /*777 QObject **/) *QDirModel {
	return NewQDirModel1(parent)
}
func NewQDirModel1(parent qtcore.QObject_ITF /*777 QObject **/) *QDirModel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDirModel")
	return gothis
}

// /usr/include/qt/QtWidgets/qdirmodel.h:70
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QDirModel(QObject *)

/*
Constructs a new directory model with the given parent. Only those files matching the nameFilters and the filters are included in the model. The sort order is given by the sort flags.
*/
func (*QDirModel) NewForInherit1p() *QDirModel {
	return NewQDirModel1p()
}
func NewQDirModel1p() *QDirModel {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDirModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDirModel")
	return gothis
}

// /usr/include/qt/QtWidgets/qdirmodel.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDirModel()

/*

 */
func DeleteQDirModel(this *QDirModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::index().

Returns the model item index for the item in the parent with the given row and column.
*/
func (this *QDirModel) Index(row int, column int, parent qtcore.QModelIndex_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::index().

Returns the model item index for the item in the parent with the given row and column.
*/
func (this *QDirModel) Indexp(row int, column int) *qtcore.QModelIndex /*123*/ {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:118
// index:1
// Public Visibility=Default Availability=Available
// [24] QModelIndex index(const QString &, int) const

/*
Reimplemented from QAbstractItemModel::index().

Returns the model item index for the item in the parent with the given row and column.
*/
func (this *QDirModel) Index1(path string, column int) *qtcore.QModelIndex /*123*/ {
	var tmpArg0 = qtcore.NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel5indexERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:118
// index:1
// Public Visibility=Default Availability=Available
// [24] QModelIndex index(const QString &, int) const

/*
Reimplemented from QAbstractItemModel::index().

Returns the model item index for the item in the parent with the given row and column.
*/
func (this *QDirModel) Index1p(path string) *qtcore.QModelIndex /*123*/ {
	var tmpArg0 = qtcore.NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, int=Int, =Invalid, , Invalid
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel5indexERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, column)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex parent(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::parent().

Return the parent of the given child model item.
*/
func (this *QDirModel) Parent(child qtcore.QModelIndex_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if child != nil && child.QModelIndex_PTR() != nil {
		convArg0 = child.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel6parentERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::rowCount().

Returns the number of rows in the parent model item.
*/
func (this *QDirModel) RowCount(parent qtcore.QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdirmodel.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::rowCount().

Returns the number of rows in the parent model item.
*/
func (this *QDirModel) RowCountp() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg0 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdirmodel.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::columnCount().

Returns the number of columns in the parent model item.
*/
func (this *QDirModel) ColumnCount(parent qtcore.QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdirmodel.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::columnCount().

Returns the number of columns in the parent model item.
*/
func (this *QDirModel) ColumnCountp() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg0 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdirmodel.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const

/*
Reimplemented from QAbstractItemModel::data().

Returns the data for the model item index with the given role.

See also setData().
*/
func (this *QDirModel) Data(index qtcore.QModelIndex_ITF, role int) *qtcore.QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const

/*
Reimplemented from QAbstractItemModel::data().

Returns the data for the model item index with the given role.

See also setData().
*/
func (this *QDirModel) Datap(index qtcore.QModelIndex_ITF) *qtcore.QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)

/*
Reimplemented from QAbstractItemModel::setData().

Sets the data for the model item index with the given role to the data referenced by the value. Returns true if successful; otherwise returns false.

See also data() and Qt::ItemDataRole.
*/
func (this *QDirModel) SetData(index qtcore.QModelIndex_ITF, value qtcore.QVariant_ITF, role int) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)

/*
Reimplemented from QAbstractItemModel::setData().

Sets the data for the model item index with the given role to the data referenced by the value. Returns true if successful; otherwise returns false.

See also data() and Qt::ItemDataRole.
*/
func (this *QDirModel) SetDatap(index qtcore.QModelIndex_ITF, value qtcore.QVariant_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::EditRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int) const

/*
Reimplemented from QAbstractItemModel::headerData().

Returns the data stored under the given role for the specified section of the header with the given orientation.
*/
func (this *QDirModel) HeaderData(section int, orientation int, role int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int) const

/*
Reimplemented from QAbstractItemModel::headerData().

Returns the data stored under the given role for the specified section of the header with the given orientation.
*/
func (this *QDirModel) HeaderDatap(section int, orientation int) *qtcore.QVariant /*123*/ {
	// arg: 2, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::hasChildren().

Returns true if the parent model item has children; otherwise returns false.
*/
func (this *QDirModel) HasChildren(index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::hasChildren().

Returns true if the parent model item has children; otherwise returns false.
*/
func (this *QDirModel) HasChildrenp() bool {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg0 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::flags().

Returns the item flags for the given index in the model.

See also Qt::ItemFlags.
*/
func (this *QDirModel) Flags(index qtcore.QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)

/*
Reimplemented from QAbstractItemModel::sort().

Sort the model items in the column using the order given. The order is a value defined in Qt::SortOrder.
*/
func (this *QDirModel) Sort(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)

/*
Reimplemented from QAbstractItemModel::sort().

Sort the model items in the column using the order given. The order is a value defined in Qt::SortOrder.
*/
func (this *QDirModel) Sortp(column int) {
	// arg: 1, Qt::SortOrder=Elaborated, Qt::SortOrder=Enum, , Invalid
	order := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStringList mimeTypes() const

/*
Reimplemented from QAbstractItemModel::mimeTypes().

Returns a list of MIME types that can be used to describe a list of items in the model.
*/
func (this *QDirModel) MimeTypes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel9mimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::dropMimeData().

Handles the data supplied by a drag and drop operation that ended with the given action over the row in the model specified by the row and column and by the parent index.

Returns true if the drop was successful, and false otherwise.

See also supportedDropActions().
*/
func (this *QDirModel) DropMimeData(data qtcore.QMimeData_ITF /*777 const QMimeData **/, action int, row int, column int, parent qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg4 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:93
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions() const

/*
Reimplemented from QAbstractItemModel::supportedDropActions().

Returns the drop actions supported by this model.

See also Qt::DropActions.
*/
func (this *QDirModel) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconProvider(QFileIconProvider *)

/*
Sets the provider of file icons for the directory model.

See also iconProvider().
*/
func (this *QDirModel) SetIconProvider(provider QFileIconProvider_ITF /*777 QFileIconProvider **/) {
	var convArg0 unsafe.Pointer
	if provider != nil && provider.QFileIconProvider_PTR() != nil {
		convArg0 = provider.QFileIconProvider_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel15setIconProviderEP17QFileIconProvider", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QFileIconProvider * iconProvider() const

/*
Returns the file icon provider for this directory model.

See also setIconProvider().
*/
func (this *QDirModel) IconProvider() *QFileIconProvider /*777 QFileIconProvider **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel12iconProviderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQFileIconProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdirmodel.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNameFilters(const QStringList &)

/*
Sets the name filters for the directory model.

See also nameFilters().
*/
func (this *QDirModel) SetNameFilters(filters qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if filters != nil && filters.QStringList_PTR() != nil {
		convArg0 = filters.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel14setNameFiltersERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList nameFilters() const

/*
Returns a list of filters applied to the names in the model.

See also setNameFilters().
*/
func (this *QDirModel) NameFilters() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel11nameFiltersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilter(QDir::Filters)

/*
Sets the directory model's filter to that specified by filters.

Note that the filter you set should always include the QDir::AllDirs enum value, otherwise QDirModel won't be able to read the directory structure.

See also filter() and QDir::Filters.
*/
func (this *QDirModel) SetFilter(filters int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel9setFilterE6QFlagsIN4QDir6FilterEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filters)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:104
// index:0
// Public Visibility=Default Availability=Available
// [4] QDir::Filters filter() const

/*
Returns the filter specification for the directory model.

See also setFilter() and QDir::Filters.
*/
func (this *QDirModel) Filter() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel6filterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSorting(QDir::SortFlags)

/*
Sets the directory model's sorting order to that specified by sort.

See also sorting() and QDir::SortFlags.
*/
func (this *QDirModel) SetSorting(sort int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel10setSortingE6QFlagsIN4QDir8SortFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), sort)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] QDir::SortFlags sorting() const

/*
Returns the sorting method used for the directory model.

See also setSorting() and QDir::SortFlags.
*/
func (this *QDirModel) Sorting() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel7sortingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResolveSymlinks(bool)

/*

 */
func (this *QDirModel) SetResolveSymlinks(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel18setResolveSymlinksEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:110
// index:0
// Public Visibility=Default Availability=Available
// [1] bool resolveSymlinks() const

/*

 */
func (this *QDirModel) ResolveSymlinks() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel15resolveSymlinksEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setReadOnly(bool)

/*

 */
func (this *QDirModel) SetReadOnly(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel11setReadOnlyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isReadOnly() const

/*

 */
func (this *QDirModel) IsReadOnly() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel10isReadOnlyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLazyChildCount(bool)

/*

 */
func (this *QDirModel) SetLazyChildCount(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel17setLazyChildCountEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool lazyChildCount() const

/*

 */
func (this *QDirModel) LazyChildCount() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel14lazyChildCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDir(const QModelIndex &) const

/*
Returns true if the model item index represents a directory; otherwise returns false.
*/
func (this *QDirModel) IsDir(index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel5isDirERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:121
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex mkdir(const QModelIndex &, const QString &)

/*
Create a directory with the name in the parent model item.
*/
func (this *QDirModel) Mkdir(parent qtcore.QModelIndex_ITF, name string) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(name)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel5mkdirERK11QModelIndexRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:122
// index:0
// Public Visibility=Default Availability=Available
// [1] bool rmdir(const QModelIndex &)

/*
Removes the directory corresponding to the model item index in the directory model and deletes the corresponding directory from the file system, returning true if successful. If the directory cannot be removed, false is returned.

Warning: This function deletes directories from the file system; it does not move them to a location where they can be recovered.

See also remove().
*/
func (this *QDirModel) Rmdir(index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel5rmdirERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:123
// index:0
// Public Visibility=Default Availability=Available
// [1] bool remove(const QModelIndex &)

/*
Removes the model item index from the directory model and deletes the corresponding file from the file system, returning true if successful. If the item cannot be removed, false is returned.

Warning: This function deletes files from the file system; it does not move them to a location where they can be recovered.

See also rmdir().
*/
func (this *QDirModel) Remove(index qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel6removeERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdirmodel.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath(const QModelIndex &) const

/*
Returns the path of the item stored in the model under the index given.
*/
func (this *QDirModel) FilePath(index qtcore.QModelIndex_ITF) string {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel8filePathERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qdirmodel.h:126
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName(const QModelIndex &) const

/*
Returns the name of the item stored in the model under the index given.
*/
func (this *QDirModel) FileName(index qtcore.QModelIndex_ITF) string {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel8fileNameERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qdirmodel.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon fileIcon(const QModelIndex &) const

/*
Returns the icons for the item stored in the model under the given index.
*/
func (this *QDirModel) FileIcon(index qtcore.QModelIndex_ITF) *qtgui.QIcon /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel8fileIconERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:128
// index:0
// Public Visibility=Default Availability=Available
// [8] QFileInfo fileInfo(const QModelIndex &) const

/*
Returns the file information for the specified model index.

Note: If the model index represents a symbolic link in the underlying filing system, the file information returned will contain information about the symbolic link itself, regardless of whether resolveSymlinks is enabled or not.

See also QFileInfo::symLinkTarget().
*/
func (this *QDirModel) FileInfo(index qtcore.QModelIndex_ITF) *qtcore.QFileInfo /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QDirModel8fileInfoERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQFileInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQFileInfo)
	return rv2
}

// /usr/include/qt/QtWidgets/qdirmodel.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refresh(const QModelIndex &)

/*
QDirModel caches file information. This function updates the cache. The parent parameter is the directory from which the model is updated; the default value will update the model from root directory of the file system (the entire model).
*/
func (this *QDirModel) Refresh(parent qtcore.QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel7refreshERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdirmodel.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refresh(const QModelIndex &)

/*
QDirModel caches file information. This function updates the cache. The parent parameter is the directory from which the model is updated; the default value will update the model from root directory of the file system (the entire model).
*/
func (this *QDirModel) Refreshp() {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg0 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QDirModel7refreshERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
ConstantValue
QDirModel::FileIconRoleQt::DecorationRole
QDirModel::FilePathRoleQt::UserRole + 1

*/
type QDirModel__Roles = int

//
const QDirModel__FileIconRole QDirModel__Roles = 1

//
const QDirModel__FilePathRole QDirModel__Roles = 257

//
const QDirModel__FileNameRole QDirModel__Roles = 258

func (this *QDirModel) RolesItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QDirModel_RolesItemName(val int) string {
	var nilthis *QDirModel
	return nilthis.RolesItemName(val)
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
