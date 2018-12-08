package qtcore

// /usr/include/qt/QtCore/qstringlistmodel.h
// #include <qstringlistmodel.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QStringListModel struct {
	*QAbstractListModel
}
type QStringListModel_ITF interface {
	QAbstractListModel_ITF
	QStringListModel_PTR() *QStringListModel
}

func (ptr *QStringListModel) QStringListModel_PTR() *QStringListModel { return ptr }

func (this *QStringListModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractListModel.GetCthis()
	}
}
func (this *QStringListModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractListModel = NewQAbstractListModelFromPointer(cthis)
}
func NewQStringListModelFromPointer(cthis unsafe.Pointer) *QStringListModel {
	bcthis0 := NewQAbstractListModelFromPointer(cthis)
	return &QStringListModel{bcthis0}
}
func (*QStringListModel) NewFromPointer(cthis unsafe.Pointer) *QStringListModel {
	return NewQStringListModelFromPointer(cthis)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QStringListModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QStringListModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qstringlistmodel.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStringListModel(QObject *)

/*
Constructs a string list model with the given parent.
*/
func (*QStringListModel) NewForInherit(parent QObject_ITF /*777 QObject **/) *QStringListModel {
	return NewQStringListModel(parent)
}
func NewQStringListModel(parent QObject_ITF /*777 QObject **/) *QStringListModel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringListModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStringListModel")
	return gothis
}

// /usr/include/qt/QtCore/qstringlistmodel.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStringListModel(QObject *)

/*
Constructs a string list model with the given parent.
*/
func (*QStringListModel) NewForInheritp() *QStringListModel {
	return NewQStringListModelp()
}
func NewQStringListModelp() *QStringListModel {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringListModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStringListModel")
	return gothis
}

// /usr/include/qt/QtCore/qstringlistmodel.h:56
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStringListModel(const QStringList &, QObject *)

/*
Constructs a string list model with the given parent.
*/
func (*QStringListModel) NewForInherit1(strings QStringList_ITF, parent QObject_ITF /*777 QObject **/) *QStringListModel {
	return NewQStringListModel1(strings, parent)
}
func NewQStringListModel1(strings QStringList_ITF, parent QObject_ITF /*777 QObject **/) *QStringListModel {
	var convArg0 unsafe.Pointer
	if strings != nil && strings.QStringList_PTR() != nil {
		convArg0 = strings.QStringList_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModelC2ERK11QStringListP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringListModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStringListModel")
	return gothis
}

// /usr/include/qt/QtCore/qstringlistmodel.h:56
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStringListModel(const QStringList &, QObject *)

/*
Constructs a string list model with the given parent.
*/
func (*QStringListModel) NewForInherit1p(strings QStringList_ITF) *QStringListModel {
	return NewQStringListModel1p(strings)
}
func NewQStringListModel1p(strings QStringList_ITF) *QStringListModel {
	var convArg0 unsafe.Pointer
	if strings != nil && strings.QStringList_PTR() != nil {
		convArg0 = strings.QStringList_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModelC2ERK11QStringListP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStringListModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStringListModel")
	return gothis
}

// /usr/include/qt/QtCore/qstringlistmodel.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::rowCount().

Returns the number of rows in the model. This value corresponds to the number of items in the model's internal string list.

The optional parent argument is in most models used to specify the parent of the rows to be counted. Because this is a list if a valid parent is specified, the result will always be 0.

See also insertRows(), removeRows(), and QAbstractItemModel::rowCount().
*/
func (this *QStringListModel) RowCount(parent QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QStringListModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlistmodel.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::rowCount().

Returns the number of rows in the model. This value corresponds to the number of items in the model's internal string list.

The optional parent argument is in most models used to specify the parent of the rows to be counted. Because this is a list if a valid parent is specified, the result will always be 0.

See also insertRows(), removeRows(), and QAbstractItemModel::rowCount().
*/
func (this *QStringListModel) RowCountp() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QStringListModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qstringlistmodel.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::sibling().
*/
func (this *QStringListModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if idx != nil && idx.QModelIndex_PTR() != nil {
		convArg2 = idx.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QStringListModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qstringlistmodel.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const

/*
Reimplemented from QAbstractItemModel::data().

Returns data for the specified role, from the item with the given index.

If the view requests an invalid index, an invalid variant is returned.

See also setData().
*/
func (this *QStringListModel) Data(index QModelIndex_ITF, role int) *QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QStringListModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qstringlistmodel.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const

/*
Reimplemented from QAbstractItemModel::data().

Returns data for the specified role, from the item with the given index.

If the view requests an invalid index, an invalid variant is returned.

See also setData().
*/
func (this *QStringListModel) Datap(index QModelIndex_ITF) *QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QStringListModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qstringlistmodel.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)

/*
Reimplemented from QAbstractItemModel::setData().

Sets the data for the specified role in the item with the given index in the model, to the provided value.

The dataChanged() signal is emitted if the item is changed.

See also Qt::ItemDataRole and data().
*/
func (this *QStringListModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlistmodel.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)

/*
Reimplemented from QAbstractItemModel::setData().

Sets the data for the specified role in the item with the given index in the model, to the provided value.

The dataChanged() signal is emitted if the item is changed.

See also Qt::ItemDataRole and data().
*/
func (this *QStringListModel) SetDatap(index QModelIndex_ITF, value QVariant_ITF) bool {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlistmodel.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::flags().

Returns the flags for the item with the given index.

Valid items are enabled, selectable, editable, drag enabled and drop enabled.

See also QAbstractItemModel::flags().
*/
func (this *QStringListModel) Flags(index QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QStringListModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::insertRows().

Inserts count rows into the model, beginning at the given row.

The parent index of the rows is optional and is only used for consistency with QAbstractItemModel. By default, a null index is specified, indicating that the rows are inserted in the top level of the model.

See also QAbstractItemModel::insertRows().
*/
func (this *QStringListModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlistmodel.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::insertRows().

Inserts count rows into the model, beginning at the given row.

The parent index of the rows is optional and is only used for consistency with QAbstractItemModel. By default, a null index is specified, indicating that the rows are inserted in the top level of the model.

See also QAbstractItemModel::insertRows().
*/
func (this *QStringListModel) InsertRowsp(row int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlistmodel.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::removeRows().

Removes count rows from the model, beginning at the given row.

The parent index of the rows is optional and is only used for consistency with QAbstractItemModel. By default, a null index is specified, indicating that the rows are removed in the top level of the model.

See also QAbstractItemModel::removeRows().
*/
func (this *QStringListModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlistmodel.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::removeRows().

Removes count rows from the model, beginning at the given row.

The parent index of the rows is optional and is only used for consistency with QAbstractItemModel. By default, a null index is specified, indicating that the rows are removed in the top level of the model.

See also QAbstractItemModel::removeRows().
*/
func (this *QStringListModel) RemoveRowsp(row int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qstringlistmodel.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)

/*
Reimplemented from QAbstractItemModel::sort().
*/
func (this *QStringListModel) Sort(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)

/*
Reimplemented from QAbstractItemModel::sort().
*/
func (this *QStringListModel) Sortp(column int) {
	// arg: 1, Qt::SortOrder=Elaborated, Qt::SortOrder=Enum, , Invalid
	order := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList stringList() const

/*
Returns the string list used by the model to store data.

See also setStringList().
*/
func (this *QStringListModel) StringList() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QStringListModel10stringListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qstringlistmodel.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStringList(const QStringList &)

/*
Sets the model's internal string list to strings. The model will notify any attached views that its underlying data has changed.

See also stringList() and dataChanged().
*/
func (this *QStringListModel) SetStringList(strings QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if strings != nil && strings.QStringList_PTR() != nil {
		convArg0 = strings.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModel13setStringListERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qstringlistmodel.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions() const

/*
Reimplemented from QAbstractItemModel::supportedDropActions().
*/
func (this *QStringListModel) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QStringListModel20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

func DeleteQStringListModel(this *QStringListModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QStringListModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
}

//  keep block end
