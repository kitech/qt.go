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
// extern C begin: 15
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
type QPersistentModelIndex struct {
	*qtrt.CObject
}
type QPersistentModelIndex_ITF interface {
	QPersistentModelIndex_PTR() *QPersistentModelIndex
}

func (ptr *QPersistentModelIndex) QPersistentModelIndex_PTR() *QPersistentModelIndex { return ptr }

func (this *QPersistentModelIndex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPersistentModelIndex) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPersistentModelIndexFromPointer(cthis unsafe.Pointer) *QPersistentModelIndex {
	return &QPersistentModelIndex{&qtrt.CObject{cthis}}
}
func (*QPersistentModelIndex) NewFromPointer(cthis unsafe.Pointer) *QPersistentModelIndex {
	return NewQPersistentModelIndexFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPersistentModelIndex()

/*

 */
func (*QPersistentModelIndex) NewForInherit() *QPersistentModelIndex {
	return NewQPersistentModelIndex()
}
func NewQPersistentModelIndex() *QPersistentModelIndex {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPersistentModelIndexC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPersistentModelIndex)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:108
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPersistentModelIndex(const QModelIndex &)

/*

 */
func (*QPersistentModelIndex) NewForInherit_1(index QModelIndex_ITF) *QPersistentModelIndex {
	return NewQPersistentModelIndex_1(index)
}
func NewQPersistentModelIndex_1(index QModelIndex_ITF) *QPersistentModelIndex {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPersistentModelIndexC2ERK11QModelIndex", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPersistentModelIndex)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPersistentModelIndex()

/*

 */
func DeleteQPersistentModelIndex(this *QPersistentModelIndex) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPersistentModelIndexD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:111
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator<(const QPersistentModelIndex &) const

/*

 */
func (this *QPersistentModelIndex) Operator_less_than(other QPersistentModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPersistentModelIndex_PTR() != nil {
		convArg0 = other.QPersistentModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndexltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QPersistentModelIndex &) const

/*

 */
func (this *QPersistentModelIndex) Operator_equal_equal(other QPersistentModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPersistentModelIndex_PTR() != nil {
		convArg0 = other.QPersistentModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndexeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:123
// index:1
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QModelIndex &) const

/*

 */
func (this *QPersistentModelIndex) Operator_equal_equal_1(other QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QModelIndex_PTR() != nil {
		convArg0 = other.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndexeqERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QPersistentModelIndex &) const

/*

 */
func (this *QPersistentModelIndex) Operator_not_equal(other QPersistentModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPersistentModelIndex_PTR() != nil {
		convArg0 = other.QPersistentModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndexneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:124
// index:1
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QModelIndex &) const

/*

 */
func (this *QPersistentModelIndex) Operator_not_equal_1(other QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QModelIndex_PTR() != nil {
		convArg0 = other.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndexneERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] QPersistentModelIndex & operator=(const QPersistentModelIndex &)

/*

 */
func (this *QPersistentModelIndex) Operator_equal(other QPersistentModelIndex_ITF) *QPersistentModelIndex {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPersistentModelIndex_PTR() != nil {
		convArg0 = other.QPersistentModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPersistentModelIndexaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPersistentModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:119
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QPersistentModelIndex & operator=(QPersistentModelIndex &&)

/*

 */
func (this *QPersistentModelIndex) Operator_equal_1(other unsafe.Pointer /*333*/) *QPersistentModelIndex {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPersistentModelIndexaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPersistentModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:125
// index:2
// Public Visibility=Default Availability=Available
// [8] QPersistentModelIndex & operator=(const QModelIndex &)

/*

 */
func (this *QPersistentModelIndex) Operator_equal_2(other QModelIndex_ITF) *QPersistentModelIndex {
	var convArg0 unsafe.Pointer
	if other != nil && other.QModelIndex_PTR() != nil {
		convArg0 = other.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPersistentModelIndexaSERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPersistentModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPersistentModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QPersistentModelIndex &)

/*

 */
func (this *QPersistentModelIndex) Swap(other QPersistentModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QPersistentModelIndex_PTR() != nil {
		convArg0 = other.QPersistentModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QPersistentModelIndex4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:127
// index:0
// Public Visibility=Default Availability=Available
// [4] int row() const

/*

 */
func (this *QPersistentModelIndex) Row() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex3rowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:128
// index:0
// Public Visibility=Default Availability=Available
// [4] int column() const

/*

 */
func (this *QPersistentModelIndex) Column() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex6columnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] void * internalPointer() const

/*

 */
func (this *QPersistentModelIndex) InternalPointer() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex15internalPointerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] quintptr internalId() const

/*

 */
func (this *QPersistentModelIndex) InternalId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex10internalIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:131
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex parent() const

/*
Returns the parent of the model item with the given index. If the item has no parent, an invalid QModelIndex is returned.

A common convention used in models that expose tree data structures is that only items in the first column have children. For that case, when reimplementing this function in a subclass the column of the returned QModelIndex would be 0.

When reimplementing this function in a subclass, be careful to avoid calling QModelIndex member functions, such as QModelIndex::parent(), since indexes belonging to your model will simply call your implementation, leading to infinite recursion.

See also createIndex().
*/
func (this *QPersistentModelIndex) Parent() *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:132
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int) const

/*
Returns the sibling at row and column for the item at index, or an invalid QModelIndex if there is no sibling at that location.

sibling() is just a convenience function that finds the item's parent, and uses it to retrieve the index of the child item in the specified row and column.

This method can optionally be overridden for implementation-specific optimization.

See also index(), QModelIndex::row(), and QModelIndex::column().
*/
func (this *QPersistentModelIndex) Sibling(row int, column int) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex7siblingEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:134
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex child(int, int) const

/*

 */
func (this *QPersistentModelIndex) Child(row int, column int) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex5childEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:136
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant data(int) const

/*
Returns the data stored under the given role for the item referred to by the index.

Note: If you do not have a value to return, return an invalid QVariant instead of returning 0.

See also Qt::ItemDataRole, setData(), and headerData().
*/
func (this *QPersistentModelIndex) Data(role int) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex4dataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:136
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant data(int) const

/*
Returns the data stored under the given role for the item referred to by the index.

Note: If you do not have a value to return, return an invalid QVariant instead of returning 0.

See also Qt::ItemDataRole, setData(), and headerData().
*/
func (this *QPersistentModelIndex) Data__() *QVariant /*123*/ {
	// arg: 0, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex4dataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags() const

/*
Returns the item flags for the given index.

The base class implementation returns a combination of flags that enables the item (ItemIsEnabled) and allows it to be selected (ItemIsSelectable).

See also Qt::ItemFlags.
*/
func (this *QPersistentModelIndex) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] const QAbstractItemModel * model() const

/*

 */
func (this *QPersistentModelIndex) Model() *QAbstractItemModel /*777 const QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:139
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QPersistentModelIndex) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QPersistentModelIndex7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
