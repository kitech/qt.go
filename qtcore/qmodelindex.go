// +build !minimal

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
// extern C begin: 5
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
type QModelIndex struct {
	*qtrt.CObject
}
type QModelIndex_ITF interface {
	QModelIndex_PTR() *QModelIndex
}

func (ptr *QModelIndex) QModelIndex_PTR() *QModelIndex { return ptr }

func (this *QModelIndex) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QModelIndex) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQModelIndexFromPointer(cthis unsafe.Pointer) *QModelIndex {
	return &QModelIndex{&qtrt.CObject{cthis}}
}
func (*QModelIndex) NewFromPointer(cthis unsafe.Pointer) *QModelIndex {
	return NewQModelIndexFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QModelIndex()

/*

 */
func (*QModelIndex) NewForInherit() *QModelIndex {
	return NewQModelIndex()
}
func NewQModelIndex() *QModelIndex {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QModelIndexC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQModelIndex)
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:62
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int row() const

/*

 */
func (this *QModelIndex) Row() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex3rowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int column() const

/*

 */
func (this *QModelIndex) Column() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex6columnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [8] quintptr internalId() const

/*

 */
func (this *QModelIndex) InternalId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex10internalIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:65
// index:0
// Public inline Visibility=Default Availability=Available
// [8] void * internalPointer() const

/*

 */
func (this *QModelIndex) InternalPointer() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex15internalPointerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:66
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex parent() const

/*
Returns the parent of the model item with the given index. If the item has no parent, an invalid QModelIndex is returned.

A common convention used in models that expose tree data structures is that only items in the first column have children. For that case, when reimplementing this function in a subclass the column of the returned QModelIndex would be 0.

When reimplementing this function in a subclass, be careful to avoid calling QModelIndex member functions, such as QModelIndex::parent(), since indexes belonging to your model will simply call your implementation, leading to infinite recursion.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also createIndex().
*/
func (this *QModelIndex) Parent() *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:67
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int) const

/*
Returns the sibling at row and column for the item at index, or an invalid QModelIndex if there is no sibling at that location.

sibling() is just a convenience function that finds the item's parent, and uses it to retrieve the index of the child item in the specified row and column.

This method can optionally be overridden for implementation-specific optimization.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also index(), QModelIndex::row(), and QModelIndex::column().
*/
func (this *QModelIndex) Sibling(row int, column int) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex7siblingEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:68
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex siblingAtColumn(int) const

/*

 */
func (this *QModelIndex) SiblingAtColumn(column int) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex15siblingAtColumnEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex siblingAtRow(int) const

/*

 */
func (this *QModelIndex) SiblingAtRow(row int) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex12siblingAtRowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [24] QModelIndex child(int, int) const

/*

 */
func (this *QModelIndex) Child(row int, column int) *QModelIndex /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex5childEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QVariant data(int) const

/*
Returns the data stored under the given role for the item referred to by the index.

Note: If you do not have a value to return, return an invalid QVariant instead of returning 0.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also Qt::ItemDataRole, setData(), and headerData().
*/
func (this *QModelIndex) Data(role int) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex4dataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:73
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QVariant data(int) const

/*
Returns the data stored under the given role for the item referred to by the index.

Note: If you do not have a value to return, return an invalid QVariant instead of returning 0.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also Qt::ItemDataRole, setData(), and headerData().
*/
func (this *QModelIndex) Datap() *QVariant /*123*/ {
	// arg: 0, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex4dataEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:74
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags() const

/*
Returns the item flags for the given index.

The base class implementation returns a combination of flags that enables the item (ItemIsEnabled) and allows it to be selected (ItemIsSelectable).

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also Qt::ItemFlags.
*/
func (this *QModelIndex) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:75
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QAbstractItemModel * model() const

/*

 */
func (this *QModelIndex) Model() *QAbstractItemModel /*777 const QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*

 */
func (this *QModelIndex) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndex7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QModelIndex &) const

/*

 */
func (this *QModelIndex) Operator_equal_equal(other QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QModelIndex_PTR() != nil {
		convArg0 = other.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndexeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QModelIndex &) const

/*

 */
func (this *QModelIndex) Operator_not_equal(other QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QModelIndex_PTR() != nil {
		convArg0 = other.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndexneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator<(const QModelIndex &) const

/*

 */
func (this *QModelIndex) Operator_less_than(other QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QModelIndex_PTR() != nil {
		convArg0 = other.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QModelIndexltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

func DeleteQModelIndex(this *QModelIndex) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QModelIndexD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10283() {
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
