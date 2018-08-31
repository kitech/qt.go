package qtcore

// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// void resetInternalData()
func (this *QAbstractItemModel) InheritResetInternalData(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "resetInternalData", f)
}

// QModelIndex createIndex(int, int, void *)
func (this *QAbstractItemModel) InheritCreateIndex(f func(row int, column int, data unsafe.Pointer /*666*/) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "createIndex", f)
}

// bool decodeData(int, int, const QModelIndex &, QDataStream &)
func (this *QAbstractItemModel) InheritDecodeData(f func(row int, column int, parent *QModelIndex, stream *QDataStream) bool) {
	qtrt.SetAllInheritCallback(this, "decodeData", f)
}

// void beginInsertRows(const QModelIndex &, int, int)
func (this *QAbstractItemModel) InheritBeginInsertRows(f func(parent *QModelIndex, first int, last int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "beginInsertRows", f)
}

// void endInsertRows()
func (this *QAbstractItemModel) InheritEndInsertRows(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "endInsertRows", f)
}

// void beginRemoveRows(const QModelIndex &, int, int)
func (this *QAbstractItemModel) InheritBeginRemoveRows(f func(parent *QModelIndex, first int, last int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "beginRemoveRows", f)
}

// void endRemoveRows()
func (this *QAbstractItemModel) InheritEndRemoveRows(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "endRemoveRows", f)
}

// bool beginMoveRows(const QModelIndex &, int, int, const QModelIndex &, int)
func (this *QAbstractItemModel) InheritBeginMoveRows(f func(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool) {
	qtrt.SetAllInheritCallback(this, "beginMoveRows", f)
}

// void endMoveRows()
func (this *QAbstractItemModel) InheritEndMoveRows(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "endMoveRows", f)
}

// void beginInsertColumns(const QModelIndex &, int, int)
func (this *QAbstractItemModel) InheritBeginInsertColumns(f func(parent *QModelIndex, first int, last int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "beginInsertColumns", f)
}

// void endInsertColumns()
func (this *QAbstractItemModel) InheritEndInsertColumns(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "endInsertColumns", f)
}

// void beginRemoveColumns(const QModelIndex &, int, int)
func (this *QAbstractItemModel) InheritBeginRemoveColumns(f func(parent *QModelIndex, first int, last int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "beginRemoveColumns", f)
}

// void endRemoveColumns()
func (this *QAbstractItemModel) InheritEndRemoveColumns(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "endRemoveColumns", f)
}

// bool beginMoveColumns(const QModelIndex &, int, int, const QModelIndex &, int)
func (this *QAbstractItemModel) InheritBeginMoveColumns(f func(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool) {
	qtrt.SetAllInheritCallback(this, "beginMoveColumns", f)
}

// void endMoveColumns()
func (this *QAbstractItemModel) InheritEndMoveColumns(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "endMoveColumns", f)
}

// void beginResetModel()
func (this *QAbstractItemModel) InheritBeginResetModel(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "beginResetModel", f)
}

// void endResetModel()
func (this *QAbstractItemModel) InheritEndResetModel(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "endResetModel", f)
}

// void changePersistentIndex(const QModelIndex &, const QModelIndex &)
func (this *QAbstractItemModel) InheritChangePersistentIndex(f func(from *QModelIndex, to *QModelIndex) /*void*/) {
	qtrt.SetAllInheritCallback(this, "changePersistentIndex", f)
}

// QModelIndexList persistentIndexList()
func (this *QAbstractItemModel) InheritPersistentIndexList(f func() *QModelIndexList /*9999*/) {
	qtrt.SetAllInheritCallback(this, "persistentIndexList", f)
}

/*

 */
type QAbstractItemModel struct {
	*QObject
}
type QAbstractItemModel_ITF interface {
	QObject_ITF
	QAbstractItemModel_PTR() *QAbstractItemModel
}

func (ptr *QAbstractItemModel) QAbstractItemModel_PTR() *QAbstractItemModel { return ptr }

func (this *QAbstractItemModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractItemModel) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQAbstractItemModelFromPointer(cthis unsafe.Pointer) *QAbstractItemModel {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QAbstractItemModel{bcthis0}
}
func (*QAbstractItemModel) NewFromPointer(cthis unsafe.Pointer) *QAbstractItemModel {
	return NewQAbstractItemModelFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:167
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractItemModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemModel(QObject *)

/*
Constructs an abstract item model with the given parent.
*/
func (*QAbstractItemModel) NewForInherit(parent QObject_ITF /*777 QObject **/) *QAbstractItemModel {
	return NewQAbstractItemModel(parent)
}
func NewQAbstractItemModel(parent QObject_ITF /*777 QObject **/) *QAbstractItemModel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractItemModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemModel(QObject *)

/*
Constructs an abstract item model with the given parent.
*/
func (*QAbstractItemModel) NewForInherit__() *QAbstractItemModel {
	return NewQAbstractItemModel__()
}
func NewQAbstractItemModel__() *QAbstractItemModel {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractItemModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:175
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractItemModel()

/*

 */
func DeleteQAbstractItemModel(this *QAbstractItemModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:177
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasIndex(int, int, const QModelIndex &) const

/*
Returns true if the model returns a valid QModelIndex for row and column with parent, otherwise returns false.
*/
func (this *QAbstractItemModel) HasIndex(row int, column int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:177
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasIndex(int, int, const QModelIndex &) const

/*
Returns true if the model returns a valid QModelIndex for row and column with parent, otherwise returns false.
*/
func (this *QAbstractItemModel) HasIndex__(row int, column int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:178
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
Returns the index of the item in the model specified by the given row, column and parent index.

When reimplementing this function in a subclass, call createIndex() to generate model indexes that other components can use to refer to items in your model.

See also createIndex().
*/
func (this *QAbstractItemModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:178
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
Returns the index of the item in the model specified by the given row, column and parent index.

When reimplementing this function in a subclass, call createIndex() to generate model indexes that other components can use to refer to items in your model.

See also createIndex().
*/
func (this *QAbstractItemModel) Index__(row int, column int) *QModelIndex /*123*/ {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:180
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex parent(const QModelIndex &) const

/*
Returns the parent of the model item with the given index. If the item has no parent, an invalid QModelIndex is returned.

A common convention used in models that expose tree data structures is that only items in the first column have children. For that case, when reimplementing this function in a subclass the column of the returned QModelIndex would be 0.

When reimplementing this function in a subclass, be careful to avoid calling QModelIndex member functions, such as QModelIndex::parent(), since indexes belonging to your model will simply call your implementation, leading to infinite recursion.

See also createIndex().
*/
func (this *QAbstractItemModel) Parent(child QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if child != nil && child.QModelIndex_PTR() != nil {
		convArg0 = child.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel6parentERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:182
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &) const

/*
Returns the sibling at row and column for the item at index, or an invalid QModelIndex if there is no sibling at that location.

sibling() is just a convenience function that finds the item's parent, and uses it to retrieve the index of the child item in the specified row and column.

This method can optionally be overridden for implementation-specific optimization.

See also index(), QModelIndex::row(), and QModelIndex::column().
*/
func (this *QAbstractItemModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if idx != nil && idx.QModelIndex_PTR() != nil {
		convArg2 = idx.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:183
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const

/*
Returns the number of rows under the given parent. When the parent is valid it means that rowCount is returning the number of children of parent.

Note: When implementing a table based model, rowCount() should return 0 when the parent is valid.

See also columnCount().
*/
func (this *QAbstractItemModel) RowCount(parent QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:183
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const

/*
Returns the number of rows under the given parent. When the parent is valid it means that rowCount is returning the number of children of parent.

Note: When implementing a table based model, rowCount() should return 0 when the parent is valid.

See also columnCount().
*/
func (this *QAbstractItemModel) RowCount__() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:184
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &) const

/*
Returns the number of columns for the children of the given parent.

In most subclasses, the number of columns is independent of the parent.

For example:


  int DomModel::columnCount(const QModelIndex &/-*parent*-/) const
  {
      return 3;
  }



Note: When implementing a table based model, columnCount() should return 0 when the parent is valid.

See also rowCount().
*/
func (this *QAbstractItemModel) ColumnCount(parent QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:184
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &) const

/*
Returns the number of columns for the children of the given parent.

In most subclasses, the number of columns is independent of the parent.

For example:


  int DomModel::columnCount(const QModelIndex &/-*parent*-/) const
  {
      return 3;
  }



Note: When implementing a table based model, columnCount() should return 0 when the parent is valid.

See also rowCount().
*/
func (this *QAbstractItemModel) ColumnCount__() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:185
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &) const

/*
Returns true if parent has any children; otherwise returns false.

Use rowCount() on the parent to find out the number of children.

Note that it is undefined behavior to report that a particular index hasChildren with this method if the same index has the flag Qt::ItemNeverHasChildren set.

See also parent() and index().
*/
func (this *QAbstractItemModel) HasChildren(parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:185
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &) const

/*
Returns true if parent has any children; otherwise returns false.

Use rowCount() on the parent to find out the number of children.

Note that it is undefined behavior to report that a particular index hasChildren with this method if the same index has the flag Qt::ItemNeverHasChildren set.

See also parent() and index().
*/
func (this *QAbstractItemModel) HasChildren__() bool {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:187
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const

/*
Returns the data stored under the given role for the item referred to by the index.

Note: If you do not have a value to return, return an invalid QVariant instead of returning 0.

See also Qt::ItemDataRole, setData(), and headerData().
*/
func (this *QAbstractItemModel) Data(index QModelIndex_ITF, role int) *QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:187
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const

/*
Returns the data stored under the given role for the item referred to by the index.

Note: If you do not have a value to return, return an invalid QVariant instead of returning 0.

See also Qt::ItemDataRole, setData(), and headerData().
*/
func (this *QAbstractItemModel) Data__(index QModelIndex_ITF) *QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:188
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)

/*
Sets the role data for the item at index to value.

Returns true if successful; otherwise returns false.

The dataChanged() signal should be emitted if the data was successfully set.

The base class implementation returns false. This function and data() must be reimplemented for editable models.

See also Qt::ItemDataRole, data(), and itemData().
*/
func (this *QAbstractItemModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:188
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)

/*
Sets the role data for the item at index to value.

Returns true if successful; otherwise returns false.

The dataChanged() signal should be emitted if the data was successfully set.

The base class implementation returns false. This function and data() must be reimplemented for editable models.

See also Qt::ItemDataRole, data(), and itemData().
*/
func (this *QAbstractItemModel) SetData__(index QModelIndex_ITF, value QVariant_ITF) bool {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:190
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int) const

/*
Returns the data for the given role and section in the header with the specified orientation.

For horizontal headers, the section number corresponds to the column number. Similarly, for vertical headers, the section number corresponds to the row number.

See also Qt::ItemDataRole, setHeaderData(), and QHeaderView.
*/
func (this *QAbstractItemModel) HeaderData(section int, orientation int, role int) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:190
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int) const

/*
Returns the data for the given role and section in the header with the specified orientation.

For horizontal headers, the section number corresponds to the column number. Similarly, for vertical headers, the section number corresponds to the row number.

See also Qt::ItemDataRole, setHeaderData(), and QHeaderView.
*/
func (this *QAbstractItemModel) HeaderData__(section int, orientation int) *QVariant /*123*/ {
	// arg: 2, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:192
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setHeaderData(int, Qt::Orientation, const QVariant &, int)

/*
Sets the data for the given role and section in the header with the specified orientation to the value supplied.

Returns true if the header's data was updated; otherwise returns false.

When reimplementing this function, the headerDataChanged() signal must be emitted explicitly.

See also Qt::ItemDataRole and headerData().
*/
func (this *QAbstractItemModel) SetHeaderData(section int, orientation int, value QVariant_ITF, role int) bool {
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:192
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setHeaderData(int, Qt::Orientation, const QVariant &, int)

/*
Sets the data for the given role and section in the header with the specified orientation to the value supplied.

Returns true if the header's data was updated; otherwise returns false.

When reimplementing this function, the headerDataChanged() signal must be emitted explicitly.

See also Qt::ItemDataRole and headerData().
*/
func (this *QAbstractItemModel) SetHeaderData__(section int, orientation int, value QVariant_ITF) bool {
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::EditRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:198
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStringList mimeTypes() const

/*
Returns the list of allowed MIME types. By default, the built-in models and views use an internal MIME type: application/x-qabstractitemmodeldatalist.

When implementing drag and drop support in a custom model, if you will return data in formats other than the default internal MIME type, reimplement this function to return your list of MIME types.

If you reimplement this function in your custom model, you must also reimplement the member functions that call it: mimeData() and dropMimeData().

See also mimeData() and dropMimeData().
*/
func (this *QAbstractItemModel) MimeTypes() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel9mimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:200
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canDropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &) const

/*
Returns true if a model can accept a drop of the data. This default implementation only checks if data has at least one format in the list of mimeTypes() and if action is among the model's supportedDropActions().

Reimplement this function in your custom model, if you want to test whether the data can be dropped at row, column, parent with action. If you don't need that test, it is not necessary to reimplement this function.

See also dropMimeData() and Using drag and drop with item views.
*/
func (this *QAbstractItemModel) CanDropMimeData(data QMimeData_ITF /*777 const QMimeData **/, action int, row int, column int, parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg4 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel15canDropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:202
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)

/*
Handles the data supplied by a drag and drop operation that ended with the given action.

Returns true if the data and action were handled by the model; otherwise returns false.

The specified row, column and parent indicate the location of an item in the model where the operation ended. It is the responsibility of the model to complete the action at the correct location.

For instance, a drop action on an item in a QTreeView can result in new items either being inserted as children of the item specified by row, column, and parent, or as siblings of the item.

When row and column are -1 it means that the dropped data should be considered as dropped directly on parent. Usually this will mean appending the data as child items of parent. If row and column are greater than or equal zero, it means that the drop occurred just before the specified row and column in the specified parent.

The mimeTypes() member is called to get the list of acceptable MIME types. This default implementation assumes the default implementation of mimeTypes(), which returns a single default MIME type. If you reimplement mimeTypes() in your custom model to return multiple MIME types, you must reimplement this function to make use of them.

See also supportedDropActions(), canDropMimeData(), and Using drag and drop with item views.
*/
func (this *QAbstractItemModel) DropMimeData(data QMimeData_ITF /*777 const QMimeData **/, action int, row int, column int, parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg4 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:204
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions() const

/*
Returns the drop actions supported by this model.

The default implementation returns Qt::CopyAction. Reimplement this function if you wish to support additional actions. You must also reimplement the dropMimeData() function to handle the additional operations.

This function was introduced in  Qt 4.2.

See also dropMimeData(), Qt::DropActions, and Using drag and drop with item views.
*/
func (this *QAbstractItemModel) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:206
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDragActions() const

/*
Returns the actions supported by the data in this model.

The default implementation returns supportedDropActions(). Reimplement this function if you wish to support additional actions.

supportedDragActions() is used by QAbstractItemView::startDrag() as the default values when a drag occurs.

See also setSupportedDragActions(), Qt::DropActions, and Using drag and drop with item views.
*/
func (this *QAbstractItemModel) SupportedDragActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel20supportedDragActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:212
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)

/*
Note: The base class implementation of this function does nothing and returns false.

On models that support this, inserts count rows into the model before the given row. Items in the new row will be children of the item represented by the parent model index.

If row is 0, the rows are prepended to any existing rows in the parent.

If row is rowCount(), the rows are appended to any existing rows in the parent.

If parent has no children, a single column with count rows is inserted.

Returns true if the rows were successfully inserted; otherwise returns false.

If you implement your own model, you can reimplement this function if you want to support insertions. Alternatively, you can provide your own API for altering the data. In either case, you will need to call beginInsertRows() and endInsertRows() to notify other components that the model has changed.

See also insertColumns(), removeRows(), beginInsertRows(), and endInsertRows().
*/
func (this *QAbstractItemModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:212
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)

/*
Note: The base class implementation of this function does nothing and returns false.

On models that support this, inserts count rows into the model before the given row. Items in the new row will be children of the item represented by the parent model index.

If row is 0, the rows are prepended to any existing rows in the parent.

If row is rowCount(), the rows are appended to any existing rows in the parent.

If parent has no children, a single column with count rows is inserted.

Returns true if the rows were successfully inserted; otherwise returns false.

If you implement your own model, you can reimplement this function if you want to support insertions. Alternatively, you can provide your own API for altering the data. In either case, you will need to call beginInsertRows() and endInsertRows() to notify other components that the model has changed.

See also insertColumns(), removeRows(), beginInsertRows(), and endInsertRows().
*/
func (this *QAbstractItemModel) InsertRows__(row int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:213
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)

/*
On models that support this, inserts count new columns into the model before the given column. The items in each new column will be children of the item represented by the parent model index.

If column is 0, the columns are prepended to any existing columns.

If column is columnCount(), the columns are appended to any existing columns.

If parent has no children, a single row with count columns is inserted.

Returns true if the columns were successfully inserted; otherwise returns false.

The base class implementation does nothing and returns false.

If you implement your own model, you can reimplement this function if you want to support insertions. Alternatively, you can provide your own API for altering the data.

See also insertRows(), removeColumns(), beginInsertColumns(), and endInsertColumns().
*/
func (this *QAbstractItemModel) InsertColumns(column int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:213
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)

/*
On models that support this, inserts count new columns into the model before the given column. The items in each new column will be children of the item represented by the parent model index.

If column is 0, the columns are prepended to any existing columns.

If column is columnCount(), the columns are appended to any existing columns.

If parent has no children, a single row with count columns is inserted.

Returns true if the columns were successfully inserted; otherwise returns false.

The base class implementation does nothing and returns false.

If you implement your own model, you can reimplement this function if you want to support insertions. Alternatively, you can provide your own API for altering the data.

See also insertRows(), removeColumns(), beginInsertColumns(), and endInsertColumns().
*/
func (this *QAbstractItemModel) InsertColumns__(column int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:214
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)

/*
On models that support this, removes count rows starting with the given row under parent parent from the model.

Returns true if the rows were successfully removed; otherwise returns false.

The base class implementation does nothing and returns false.

If you implement your own model, you can reimplement this function if you want to support removing. Alternatively, you can provide your own API for altering the data.

See also removeRow(), removeColumns(), insertColumns(), beginRemoveRows(), and endRemoveRows().
*/
func (this *QAbstractItemModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:214
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)

/*
On models that support this, removes count rows starting with the given row under parent parent from the model.

Returns true if the rows were successfully removed; otherwise returns false.

The base class implementation does nothing and returns false.

If you implement your own model, you can reimplement this function if you want to support removing. Alternatively, you can provide your own API for altering the data.

See also removeRow(), removeColumns(), insertColumns(), beginRemoveRows(), and endRemoveRows().
*/
func (this *QAbstractItemModel) RemoveRows__(row int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:215
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)

/*
On models that support this, removes count columns starting with the given column under parent parent from the model.

Returns true if the columns were successfully removed; otherwise returns false.

The base class implementation does nothing and returns false.

If you implement your own model, you can reimplement this function if you want to support removing. Alternatively, you can provide your own API for altering the data.

See also removeColumn(), removeRows(), insertColumns(), beginRemoveColumns(), and endRemoveColumns().
*/
func (this *QAbstractItemModel) RemoveColumns(column int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:215
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)

/*
On models that support this, removes count columns starting with the given column under parent parent from the model.

Returns true if the columns were successfully removed; otherwise returns false.

The base class implementation does nothing and returns false.

If you implement your own model, you can reimplement this function if you want to support removing. Alternatively, you can provide your own API for altering the data.

See also removeColumn(), removeRows(), insertColumns(), beginRemoveColumns(), and endRemoveColumns().
*/
func (this *QAbstractItemModel) RemoveColumns__(column int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:216
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool moveRows(const QModelIndex &, int, int, const QModelIndex &, int)

/*
On models that support this, moves count rows starting with the given sourceRow under parent sourceParent to row destinationChild under parent destinationParent.

Returns true if the rows were successfully moved; otherwise returns false.

The base class implementation does nothing and returns false.

If you implement your own model, you can reimplement this function if you want to support moving. Alternatively, you can provide your own API for altering the data.

See also beginMoveRows() and endMoveRows().
*/
func (this *QAbstractItemModel) MoveRows(sourceParent QModelIndex_ITF, sourceRow int, count int, destinationParent QModelIndex_ITF, destinationChild int) bool {
	var convArg0 unsafe.Pointer
	if sourceParent != nil && sourceParent.QModelIndex_PTR() != nil {
		convArg0 = sourceParent.QModelIndex_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if destinationParent != nil && destinationParent.QModelIndex_PTR() != nil {
		convArg3 = destinationParent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel8moveRowsERK11QModelIndexiiS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceRow, count, convArg3, destinationChild)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:218
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool moveColumns(const QModelIndex &, int, int, const QModelIndex &, int)

/*
On models that support this, moves count columns starting with the given sourceColumn under parent sourceParent to column destinationChild under parent destinationParent.

Returns true if the columns were successfully moved; otherwise returns false.

The base class implementation does nothing and returns false.

If you implement your own model, you can reimplement this function if you want to support moving. Alternatively, you can provide your own API for altering the data.

See also beginMoveColumns() and endMoveColumns().
*/
func (this *QAbstractItemModel) MoveColumns(sourceParent QModelIndex_ITF, sourceColumn int, count int, destinationParent QModelIndex_ITF, destinationChild int) bool {
	var convArg0 unsafe.Pointer
	if sourceParent != nil && sourceParent.QModelIndex_PTR() != nil {
		convArg0 = sourceParent.QModelIndex_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if destinationParent != nil && destinationParent.QModelIndex_PTR() != nil {
		convArg3 = destinationParent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel11moveColumnsERK11QModelIndexiiS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceColumn, count, convArg3, destinationChild)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool insertRow(int, const QModelIndex &)

/*
Inserts a single row before the given row in the child items of the parent specified.

Note: This function calls the virtual method insertRows.

Returns true if the row is inserted; otherwise returns false.

See also insertRows(), insertColumn(), and removeRow().
*/
func (this *QAbstractItemModel) InsertRow(row int, parent QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel9insertRowEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool insertRow(int, const QModelIndex &)

/*
Inserts a single row before the given row in the child items of the parent specified.

Note: This function calls the virtual method insertRows.

Returns true if the row is inserted; otherwise returns false.

See also insertRows(), insertColumn(), and removeRow().
*/
func (this *QAbstractItemModel) InsertRow__(row int) bool {
	// arg: 1, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg1 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel9insertRowEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool insertColumn(int, const QModelIndex &)

/*
Inserts a single column before the given column in the child items of the parent specified.

Returns true if the column is inserted; otherwise returns false.

See also insertColumns(), insertRow(), and removeColumn().
*/
func (this *QAbstractItemModel) InsertColumn(column int, parent QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:222
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool insertColumn(int, const QModelIndex &)

/*
Inserts a single column before the given column in the child items of the parent specified.

Returns true if the column is inserted; otherwise returns false.

See also insertColumns(), insertRow(), and removeColumn().
*/
func (this *QAbstractItemModel) InsertColumn__(column int) bool {
	// arg: 1, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg1 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:223
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool removeRow(int, const QModelIndex &)

/*
Removes the given row from the child items of the parent specified.

Returns true if the row is removed; otherwise returns false.

This is a convenience function that calls removeRows(). The QAbstractItemModel implementation of removeRows() does nothing.

See also removeRows(), removeColumn(), and insertRow().
*/
func (this *QAbstractItemModel) RemoveRow(row int, parent QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel9removeRowEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:223
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool removeRow(int, const QModelIndex &)

/*
Removes the given row from the child items of the parent specified.

Returns true if the row is removed; otherwise returns false.

This is a convenience function that calls removeRows(). The QAbstractItemModel implementation of removeRows() does nothing.

See also removeRows(), removeColumn(), and insertRow().
*/
func (this *QAbstractItemModel) RemoveRow__(row int) bool {
	// arg: 1, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg1 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel9removeRowEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:224
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool removeColumn(int, const QModelIndex &)

/*
Removes the given column from the child items of the parent specified.

Returns true if the column is removed; otherwise returns false.

See also removeColumns(), removeRow(), and insertColumn().
*/
func (this *QAbstractItemModel) RemoveColumn(column int, parent QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:224
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool removeColumn(int, const QModelIndex &)

/*
Removes the given column from the child items of the parent specified.

Returns true if the column is removed; otherwise returns false.

See also removeColumns(), removeRow(), and insertColumn().
*/
func (this *QAbstractItemModel) RemoveColumn__(column int) bool {
	// arg: 1, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg1 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:225
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool moveRow(const QModelIndex &, int, const QModelIndex &, int)

/*
On models that support this, moves sourceRow from sourceParent to destinationChild under destinationParent.

Returns true if the rows were successfully moved; otherwise returns false.

See also moveRows() and moveColumn().
*/
func (this *QAbstractItemModel) MoveRow(sourceParent QModelIndex_ITF, sourceRow int, destinationParent QModelIndex_ITF, destinationChild int) bool {
	var convArg0 unsafe.Pointer
	if sourceParent != nil && sourceParent.QModelIndex_PTR() != nil {
		convArg0 = sourceParent.QModelIndex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if destinationParent != nil && destinationParent.QModelIndex_PTR() != nil {
		convArg2 = destinationParent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel7moveRowERK11QModelIndexiS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceRow, convArg2, destinationChild)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:227
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool moveColumn(const QModelIndex &, int, const QModelIndex &, int)

/*
On models that support this, moves sourceColumn from sourceParent to destinationChild under destinationParent.

Returns true if the columns were successfully moved; otherwise returns false.

See also moveColumns() and moveRow().
*/
func (this *QAbstractItemModel) MoveColumn(sourceParent QModelIndex_ITF, sourceColumn int, destinationParent QModelIndex_ITF, destinationChild int) bool {
	var convArg0 unsafe.Pointer
	if sourceParent != nil && sourceParent.QModelIndex_PTR() != nil {
		convArg0 = sourceParent.QModelIndex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if destinationParent != nil && destinationParent.QModelIndex_PTR() != nil {
		convArg2 = destinationParent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel10moveColumnERK11QModelIndexiS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceColumn, convArg2, destinationChild)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:230
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fetchMore(const QModelIndex &)

/*
Fetches any available data for the items with the parent specified by the parent index.

Reimplement this if you are populating your model incrementally.

The default implementation does nothing.

See also canFetchMore().
*/
func (this *QAbstractItemModel) FetchMore(parent QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel9fetchMoreERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:231
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canFetchMore(const QModelIndex &) const

/*
Returns true if there is more data available for parent; otherwise returns false.

The default implementation always returns false.

If canFetchMore() returns true, the fetchMore() function should be called. This is the behavior of QAbstractItemView, for example.

See also fetchMore().
*/
func (this *QAbstractItemModel) CanFetchMore(parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel12canFetchMoreERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:232
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &) const

/*
Returns the item flags for the given index.

The base class implementation returns a combination of flags that enables the item (ItemIsEnabled) and allows it to be selected (ItemIsSelectable).

See also Qt::ItemFlags.
*/
func (this *QAbstractItemModel) Flags(index QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:233
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)

/*
Sorts the model by column in the given order.

The base class implementation does nothing.
*/
func (this *QAbstractItemModel) Sort(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:233
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)

/*
Sorts the model by column in the given order.

The base class implementation does nothing.
*/
func (this *QAbstractItemModel) Sort__(column int) {
	// arg: 1, Qt::SortOrder=Elaborated, Qt::SortOrder=Enum, , Invalid
	order := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:234
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex buddy(const QModelIndex &) const

/*
Returns a model index for the buddy of the item represented by index. When the user wants to edit an item, the view will call this function to check whether another item in the model should be edited instead. Then, the view will construct a delegate using the model index returned by the buddy item.

The default implementation of this function has each item as its own buddy.
*/
func (this *QAbstractItemModel) Buddy(index QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel5buddyERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:235
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QModelIndexList match(const QModelIndex &, int, const QVariant &, int, Qt::MatchFlags) const

/*
Returns a list of indexes for the items in the column of the start index where data stored under the given role matches the specified value. The way the search is performed is defined by the flags given. The list that is returned may be empty. Note also that the order of results in the list may not correspond to the order in the model, if for example a proxy model is used. The order of the results can not be relied upon.

The search begins from the start index, and continues until the number of matching data items equals hits, the search reaches the last row, or the search reaches start again - depending on whether MatchWrap is specified in flags. If you want to search for all matching items, use hits = -1.

By default, this function will perform a wrapping, string-based comparison on all items, searching for items that begin with the search term specified by value.

Note: The default implementation of this function only searches columns. Reimplement this function to include a different search behavior.
*/
func (this *QAbstractItemModel) Match(start QModelIndex_ITF, role int, value QVariant_ITF, hits int, flags int) *QModelIndexList /*667*/ {
	var convArg0 unsafe.Pointer
	if start != nil && start.QModelIndex_PTR() != nil {
		convArg0 = start.QModelIndex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel5matchERK11QModelIndexiRK8QVarianti6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, convArg2, hits, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:235
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QModelIndexList match(const QModelIndex &, int, const QVariant &, int, Qt::MatchFlags) const

/*
Returns a list of indexes for the items in the column of the start index where data stored under the given role matches the specified value. The way the search is performed is defined by the flags given. The list that is returned may be empty. Note also that the order of results in the list may not correspond to the order in the model, if for example a proxy model is used. The order of the results can not be relied upon.

The search begins from the start index, and continues until the number of matching data items equals hits, the search reaches the last row, or the search reaches start again - depending on whether MatchWrap is specified in flags. If you want to search for all matching items, use hits = -1.

By default, this function will perform a wrapping, string-based comparison on all items, searching for items that begin with the search term specified by value.

Note: The default implementation of this function only searches columns. Reimplement this function to include a different search behavior.
*/
func (this *QAbstractItemModel) Match__(start QModelIndex_ITF, role int, value QVariant_ITF) *QModelIndexList /*667*/ {
	var convArg0 unsafe.Pointer
	if start != nil && start.QModelIndex_PTR() != nil {
		convArg0 = start.QModelIndex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid, , Invalid
	hits := int(1)
	// arg: 4, Qt::MatchFlags=Elaborated, Qt::MatchFlags=Typedef, QFlags<Qt::MatchFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel5matchERK11QModelIndexiRK8QVarianti6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, convArg2, hits, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:235
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QModelIndexList match(const QModelIndex &, int, const QVariant &, int, Qt::MatchFlags) const

/*
Returns a list of indexes for the items in the column of the start index where data stored under the given role matches the specified value. The way the search is performed is defined by the flags given. The list that is returned may be empty. Note also that the order of results in the list may not correspond to the order in the model, if for example a proxy model is used. The order of the results can not be relied upon.

The search begins from the start index, and continues until the number of matching data items equals hits, the search reaches the last row, or the search reaches start again - depending on whether MatchWrap is specified in flags. If you want to search for all matching items, use hits = -1.

By default, this function will perform a wrapping, string-based comparison on all items, searching for items that begin with the search term specified by value.

Note: The default implementation of this function only searches columns. Reimplement this function to include a different search behavior.
*/
func (this *QAbstractItemModel) Match__1(start QModelIndex_ITF, role int, value QVariant_ITF, hits int) *QModelIndexList /*667*/ {
	var convArg0 unsafe.Pointer
	if start != nil && start.QModelIndex_PTR() != nil {
		convArg0 = start.QModelIndex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	// arg: 4, Qt::MatchFlags=Elaborated, Qt::MatchFlags=Typedef, QFlags<Qt::MatchFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel5matchERK11QModelIndexiRK8QVarianti6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, convArg2, hits, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:239
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize span(const QModelIndex &) const

/*
Returns the row and column span of the item represented by index.

Note: Currently, span is not used.
*/
func (this *QAbstractItemModel) Span(index QModelIndex_ITF) *QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel4spanERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:255
// index:0
// Public Visibility=Default Availability=Available
// [-2] void headerDataChanged(Qt::Orientation, int, int)

/*
This signal is emitted whenever a header is changed. The orientation indicates whether the horizontal or vertical header has changed. The sections in the header from the first to the last need to be updated.

When reimplementing the setHeaderData() function, this signal must be emitted explicitly.

If you are changing the number of columns or rows you do not need to emit this signal, but use the begin/end functions (refer to the section on subclassing in the QAbstractItemModel class description for details).

See also headerData(), setHeaderData(), and dataChanged().
*/
func (this *QAbstractItemModel) HeaderDataChanged(orientation int, first int, last int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel17headerDataChangedEN2Qt11OrientationEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation, first, last)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:281
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool submit()

/*
Lets the model know that it should submit cached information to permanent storage. This function is typically used for row editing.

Returns true if there is no error; otherwise returns false.

See also revert().
*/
func (this *QAbstractItemModel) Submit() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel6submitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:282
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void revert()

/*
Lets the model know that it should discard cached information. This function is typically used for row editing.

See also submit().
*/
func (this *QAbstractItemModel) Revert() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel6revertEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:286
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void resetInternalData()

/*
This slot is called just after the internal data of a model is cleared while it is being reset.

This slot is provided the convenience of subclasses of concrete proxy models, such as subclasses of QSortFilterProxyModel which maintain extra data.


  class CustomDataProxy : public QSortFilterProxyModel
  {
      Q_OBJECT
  public:
      CustomDataProxy(QObject *parent)
        : QSortFilterProxyModel(parent)
      {
      }

      ...

      QVariant data(const QModelIndex &index, int role)
      {
          if (role != Qt::BackgroundRole)
              return QSortFilterProxyModel::data(index, role);

          if (m_customData.contains(index.row()))
              return m_customData.value(index.row());
          return QSortFilterProxyModel::data(index, role);
      }

  private slots:
      void resetInternalData()
      {
          m_customData.clear();
      }

  private:
    QHash<int, QVariant> m_customData;
  };



Note: Due to a mistake, this slot is missing in Qt 5.0.

This function was introduced in  Qt 4.8.

See also modelAboutToBeReset() and modelReset().
*/
func (this *QAbstractItemModel) ResetInternalData() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel17resetInternalDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:291
// index:0
// Protected inline Visibility=Default Availability=Available
// [24] QModelIndex createIndex(int, int, void *) const

/*
Creates a model index for the given row and column with the internal pointer ptr.

When using a QSortFilterProxyModel, its indexes have their own internal pointer. It is not advisable to access this internal pointer outside of the model. Use the data() function instead.

This function provides a consistent interface that model subclasses must use to create model indexes.
*/
func (this *QAbstractItemModel) CreateIndex(row int, column int, data unsafe.Pointer /*666*/) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel11createIndexEiiPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, data)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:291
// index:0
// Protected inline Visibility=Default Availability=Available
// [24] QModelIndex createIndex(int, int, void *) const

/*
Creates a model index for the given row and column with the internal pointer ptr.

When using a QSortFilterProxyModel, its indexes have their own internal pointer. It is not advisable to access this internal pointer outside of the model. Use the data() function instead.

This function provides a consistent interface that model subclasses must use to create model indexes.
*/
func (this *QAbstractItemModel) CreateIndex__(row int, column int) *QModelIndex /*123*/ {
	// arg: 2, void *=Pointer, =Invalid, , Invalid
	var data unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel11createIndexEiiPv", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, data)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:292
// index:1
// Protected inline Visibility=Default Availability=Available
// [24] QModelIndex createIndex(int, int, quintptr) const

/*
Creates a model index for the given row and column with the internal pointer ptr.

When using a QSortFilterProxyModel, its indexes have their own internal pointer. It is not advisable to access this internal pointer outside of the model. Use the data() function instead.

This function provides a consistent interface that model subclasses must use to create model indexes.
*/
func (this *QAbstractItemModel) CreateIndex_1(row int, column int, id uint64) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel11createIndexEiiy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, id)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:295
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool decodeData(int, int, const QModelIndex &, QDataStream &)

/*

 */
func (this *QAbstractItemModel) DecodeData(row int, column int, parent QModelIndex_ITF, stream QDataStream_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if stream != nil && stream.QDataStream_PTR() != nil {
		convArg3 = stream.QDataStream_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel10decodeDataEiiRK11QModelIndexR11QDataStream", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:297
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void beginInsertRows(const QModelIndex &, int, int)

/*
Begins a row insertion operation.

When reimplementing insertRows() in a subclass, you must call this function before inserting data into the model's underlying data store.

The parent index corresponds to the parent into which the new rows are inserted; first and last are the row numbers that the new rows will have after they have been inserted.


 Specify the first and last row numbers for the span of rows you want to insert into an item in a model.For example, as shown in the diagram, we insert three rows before row 2, so first is 2 and last is 4:

  beginInsertRows(parent, 2, 4);


This inserts the three new rows as rows 2, 3, and 4.

To append rows, insert them after the last row.For example, as shown in the diagram, we append two rows to a collection of 4 existing rows (ending in row 3), so first is 4 and last is 5:

  beginInsertRows(parent, 4, 5);


This appends the two new rows as rows 4 and 5.



Note: This function emits the rowsAboutToBeInserted() signal which connected views (or proxies) must handle before the data is inserted. Otherwise, the views may end up in an invalid state.

See also endInsertRows().
*/
func (this *QAbstractItemModel) BeginInsertRows(parent QModelIndex_ITF, first int, last int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel15beginInsertRowsERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:298
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endInsertRows()

/*
Ends a row insertion operation.

When reimplementing insertRows() in a subclass, you must call this function after inserting data into the model's underlying data store.

See also beginInsertRows().
*/
func (this *QAbstractItemModel) EndInsertRows() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13endInsertRowsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:300
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void beginRemoveRows(const QModelIndex &, int, int)

/*
Begins a row removal operation.

When reimplementing removeRows() in a subclass, you must call this function before removing data from the model's underlying data store.

The parent index corresponds to the parent from which the new rows are removed; first and last are the row numbers of the rows to be removed.


 Specify the first and last row numbers for the span of rows you want to remove from an item in a model.For example, as shown in the diagram, we remove the two rows from row 2 to row 3, so first is 2 and last is 3:

  beginRemoveRows(parent, 2, 3);





Note: This function emits the rowsAboutToBeRemoved() signal which connected views (or proxies) must handle before the data is removed. Otherwise, the views may end up in an invalid state.

See also endRemoveRows().
*/
func (this *QAbstractItemModel) BeginRemoveRows(parent QModelIndex_ITF, first int, last int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel15beginRemoveRowsERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:301
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endRemoveRows()

/*
Ends a row removal operation.

When reimplementing removeRows() in a subclass, you must call this function after removing data from the model's underlying data store.

See also beginRemoveRows().
*/
func (this *QAbstractItemModel) EndRemoveRows() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13endRemoveRowsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:303
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool beginMoveRows(const QModelIndex &, int, int, const QModelIndex &, int)

/*
Begins a row move operation.

When reimplementing a subclass, this method simplifies moving entities in your model. This method is responsible for moving persistent indexes in the model, which you would otherwise be required to do yourself. Using beginMoveRows and endMoveRows is an alternative to emitting layoutAboutToBeChanged and layoutChanged directly along with changePersistentIndex.

The sourceParent index corresponds to the parent from which the rows are moved; sourceFirst and sourceLast are the first and last row numbers of the rows to be moved. The destinationParent index corresponds to the parent into which those rows are moved. The destinationChild is the row to which the rows will be moved. That is, the index at row sourceFirst in sourceParent will become row destinationChild in destinationParent, followed by all other rows up to sourceLast.

However, when moving rows down in the same parent (sourceParent and destinationParent are equal), the rows will be placed before the destinationChild index. That is, if you wish to move rows 0 and 1 so they will become rows 1 and 2, destinationChild should be 3. In this case, the new index for the source row i (which is between sourceFirst and sourceLast) is equal to (destinationChild-sourceLast-1+i).

Note that if sourceParent and destinationParent are the same, you must ensure that the destinationChild is not within the range of sourceFirst and sourceLast + 1. You must also ensure that you do not attempt to move a row to one of its own children or ancestors. This method returns false if either condition is true, in which case you should abort your move operation.


 Specify the first and last row numbers for the span of rows in the source parent you want to move in the model. Also specify the row in the destination parent to move the span to.For example, as shown in the diagram, we move three rows from row 2 to 4 in the source, so sourceFirst is 2 and sourceLast is 4. We move those items to above row 2 in the destination, so destinationChild is 2.

  beginMoveRows(sourceParent, 2, 4, destinationParent, 2);


This moves the three rows rows 2, 3, and 4 in the source to become 2, 3 and 4 in the destination. Other affected siblings are displaced accordingly.

To append rows to another parent, move them to after the last row.For example, as shown in the diagram, we move three rows to a collection of 6 existing rows (ending in row 5), so destinationChild is 6:

  beginMoveRows(sourceParent, 2, 4, destinationParent, 6);


This moves the target rows to the end of the target parent as 6, 7 and 8.

To move rows within the same parent, specify the row to move them to.For example, as shown in the diagram, we move one item from row 2 to row 0, so sourceFirst and sourceLast are 2 and destinationChild is 0.

  beginMoveRows(parent, 2, 2, parent, 0);


Note that other rows may be displaced accordingly. Note also that when moving items within the same parent you should not attempt invalid or no-op moves. In the above example, item 2 is at row 2 before the move, so it can not be moved to row 2 (where it is already) or row 3 (no-op as row 3 means above row 3, where it is already)

To move rows within the same parent, specify the row to move them to.For example, as shown in the diagram, we move one item from row 2 to row 4, so sourceFirst and sourceLast are 2 and destinationChild is 4.

  beginMoveRows(parent, 2, 2, parent, 4);


Note that other rows may be displaced accordingly.



This function was introduced in  Qt 4.6.

See also endMoveRows().
*/
func (this *QAbstractItemModel) BeginMoveRows(sourceParent QModelIndex_ITF, sourceFirst int, sourceLast int, destinationParent QModelIndex_ITF, destinationRow int) bool {
	var convArg0 unsafe.Pointer
	if sourceParent != nil && sourceParent.QModelIndex_PTR() != nil {
		convArg0 = sourceParent.QModelIndex_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if destinationParent != nil && destinationParent.QModelIndex_PTR() != nil {
		convArg3 = destinationParent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13beginMoveRowsERK11QModelIndexiiS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceFirst, sourceLast, convArg3, destinationRow)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:304
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endMoveRows()

/*
Ends a row move operation.

When implementing a subclass, you must call this function after moving data within the model's underlying data store.

This function was introduced in  Qt 4.6.

See also beginMoveRows().
*/
func (this *QAbstractItemModel) EndMoveRows() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel11endMoveRowsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:306
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void beginInsertColumns(const QModelIndex &, int, int)

/*
Begins a column insertion operation.

When reimplementing insertColumns() in a subclass, you must call this function before inserting data into the model's underlying data store.

The parent index corresponds to the parent into which the new columns are inserted; first and last are the column numbers of the new columns will have after they have been inserted.


 Specify the first and last column numbers for the span of columns you want to insert into an item in a model.For example, as shown in the diagram, we insert three columns before column 4, so first is 4 and last is 6:

  beginInsertColumns(parent, 4, 6);


This inserts the three new columns as columns 4, 5, and 6.

To append columns, insert them after the last column.For example, as shown in the diagram, we append three columns to a collection of six existing columns (ending in column 5), so first is 6 and last is 8:

  beginInsertColumns(parent, 6, 8);


This appends the two new columns as columns 6, 7, and 8.



Note: This function emits the columnsAboutToBeInserted() signal which connected views (or proxies) must handle before the data is inserted. Otherwise, the views may end up in an invalid state.

See also endInsertColumns().
*/
func (this *QAbstractItemModel) BeginInsertColumns(parent QModelIndex_ITF, first int, last int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel18beginInsertColumnsERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:307
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endInsertColumns()

/*
Ends a column insertion operation.

When reimplementing insertColumns() in a subclass, you must call this function after inserting data into the model's underlying data store.

See also beginInsertColumns().
*/
func (this *QAbstractItemModel) EndInsertColumns() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel16endInsertColumnsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:309
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void beginRemoveColumns(const QModelIndex &, int, int)

/*
Begins a column removal operation.

When reimplementing removeColumns() in a subclass, you must call this function before removing data from the model's underlying data store.

The parent index corresponds to the parent from which the new columns are removed; first and last are the column numbers of the first and last columns to be removed.


 Specify the first and last column numbers for the span of columns you want to remove from an item in a model.For example, as shown in the diagram, we remove the three columns from column 4 to column 6, so first is 4 and last is 6:

  beginRemoveColumns(parent, 4, 6);





Note: This function emits the columnsAboutToBeRemoved() signal which connected views (or proxies) must handle before the data is removed. Otherwise, the views may end up in an invalid state.

See also endRemoveColumns().
*/
func (this *QAbstractItemModel) BeginRemoveColumns(parent QModelIndex_ITF, first int, last int) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel18beginRemoveColumnsERK11QModelIndexii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, first, last)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:310
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endRemoveColumns()

/*
Ends a column removal operation.

When reimplementing removeColumns() in a subclass, you must call this function after removing data from the model's underlying data store.

See also beginRemoveColumns().
*/
func (this *QAbstractItemModel) EndRemoveColumns() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel16endRemoveColumnsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:312
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool beginMoveColumns(const QModelIndex &, int, int, const QModelIndex &, int)

/*
Begins a column move operation.

When reimplementing a subclass, this method simplifies moving entities in your model. This method is responsible for moving persistent indexes in the model, which you would otherwise be required to do yourself. Using beginMoveColumns and endMoveColumns is an alternative to emitting layoutAboutToBeChanged and layoutChanged directly along with changePersistentIndex.

The sourceParent index corresponds to the parent from which the columns are moved; sourceFirst and sourceLast are the first and last column numbers of the columns to be moved. The destinationParent index corresponds to the parent into which those columns are moved. The destinationChild is the column to which the columns will be moved. That is, the index at column sourceFirst in sourceParent will become column destinationChild in destinationParent, followed by all other columns up to sourceLast.

However, when moving columns down in the same parent (sourceParent and destinationParent are equal), the columns will be placed before the destinationChild index. That is, if you wish to move columns 0 and 1 so they will become columns 1 and 2, destinationChild should be 3. In this case, the new index for the source column i (which is between sourceFirst and sourceLast) is equal to (destinationChild-sourceLast-1+i).

Note that if sourceParent and destinationParent are the same, you must ensure that the destinationChild is not within the range of sourceFirst and sourceLast + 1. You must also ensure that you do not attempt to move a column to one of its own children or ancestors. This method returns false if either condition is true, in which case you should abort your move operation.

This function was introduced in  Qt 4.6.

See also endMoveColumns().
*/
func (this *QAbstractItemModel) BeginMoveColumns(sourceParent QModelIndex_ITF, sourceFirst int, sourceLast int, destinationParent QModelIndex_ITF, destinationColumn int) bool {
	var convArg0 unsafe.Pointer
	if sourceParent != nil && sourceParent.QModelIndex_PTR() != nil {
		convArg0 = sourceParent.QModelIndex_PTR().GetCthis()
	}
	var convArg3 unsafe.Pointer
	if destinationParent != nil && destinationParent.QModelIndex_PTR() != nil {
		convArg3 = destinationParent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel16beginMoveColumnsERK11QModelIndexiiS2_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, sourceFirst, sourceLast, convArg3, destinationColumn)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:313
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endMoveColumns()

/*
Ends a column move operation.

When implementing a subclass, you must call this function after moving data within the model's underlying data store.

This function was introduced in  Qt 4.6.

See also beginMoveColumns().
*/
func (this *QAbstractItemModel) EndMoveColumns() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel14endMoveColumnsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:324
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void beginResetModel()

/*
Begins a model reset operation.

A reset operation resets the model to its current state in any attached views.

Note: Any views attached to this model will be reset as well.

When a model is reset it means that any previous data reported from the model is now invalid and has to be queried for again. This also means that the current item and any selected items will become invalid.

When a model radically changes its data it can sometimes be easier to just call this function rather than emit dataChanged() to inform other components when the underlying data source, or its structure, has changed.

You must call this function before resetting any internal data structures in your model or proxy model.

This function emits the signal modelAboutToBeReset().

This function was introduced in  Qt 4.6.

See also modelAboutToBeReset(), modelReset(), and endResetModel().
*/
func (this *QAbstractItemModel) BeginResetModel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel15beginResetModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:325
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void endResetModel()

/*
Completes a model reset operation.

You must call this function after resetting any internal data structure in your model or proxy model.

This function emits the signal modelReset().

This function was introduced in  Qt 4.6.

See also beginResetModel().
*/
func (this *QAbstractItemModel) EndResetModel() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel13endResetModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:327
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void changePersistentIndex(const QModelIndex &, const QModelIndex &)

/*
Changes the QPersistentModelIndex that is equal to the given from model index to the given to model index.

If no persistent model index equal to the given from model index was found, nothing is changed.

See also persistentIndexList() and changePersistentIndexList().
*/
func (this *QAbstractItemModel) ChangePersistentIndex(from QModelIndex_ITF, to QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if from != nil && from.QModelIndex_PTR() != nil {
		convArg0 = from.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if to != nil && to.QModelIndex_PTR() != nil {
		convArg1 = to.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractItemModel21changePersistentIndexERK11QModelIndexS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:329
// index:0
// Protected Visibility=Default Availability=Available
// [8] QModelIndexList persistentIndexList() const

/*
Returns the list of indexes stored as persistent indexes in the model.

This function was introduced in  Qt 4.2.
*/
func (this *QAbstractItemModel) PersistentIndexList() *QModelIndexList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractItemModel19persistentIndexListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

/*
This enum describes the way the model changes layout.



Note that VerticalSortHint and HorizontalSortHint carry the meaning that items are being moved within the same parent, not moved to a different parent in the model, and not filtered out or in.

*/
type QAbstractItemModel__LayoutChangeHint = int

// No hint is available.
const QAbstractItemModel__NoLayoutChangeHint QAbstractItemModel__LayoutChangeHint = 0

// Rows are being sorted.
const QAbstractItemModel__VerticalSortHint QAbstractItemModel__LayoutChangeHint = 1

// Columns are being sorted.
const QAbstractItemModel__HorizontalSortHint QAbstractItemModel__LayoutChangeHint = 2

func (this *QAbstractItemModel) LayoutChangeHintItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemModel_LayoutChangeHintItemName(val int) string {
	var nilthis *QAbstractItemModel
	return nilthis.LayoutChangeHintItemName(val)
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
