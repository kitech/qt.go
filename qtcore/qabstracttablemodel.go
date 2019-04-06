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
// extern C begin: 62
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
type QAbstractTableModel struct {
	*QAbstractItemModel
}
type QAbstractTableModel_ITF interface {
	QAbstractItemModel_ITF
	QAbstractTableModel_PTR() *QAbstractTableModel
}

func (ptr *QAbstractTableModel) QAbstractTableModel_PTR() *QAbstractTableModel { return ptr }

func (this *QAbstractTableModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemModel.GetCthis()
	}
}
func (this *QAbstractTableModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemModel = NewQAbstractItemModelFromPointer(cthis)
}
func NewQAbstractTableModelFromPointer(cthis unsafe.Pointer) *QAbstractTableModel {
	bcthis0 := NewQAbstractItemModelFromPointer(cthis)
	return &QAbstractTableModel{bcthis0}
}
func (*QAbstractTableModel) NewFromPointer(cthis unsafe.Pointer) *QAbstractTableModel {
	return NewQAbstractTableModelFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:387
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractTableModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractTableModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:390
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractTableModel(QObject *)

/*

 */
func (*QAbstractTableModel) NewForInherit(parent QObject_ITF /*777 QObject **/) *QAbstractTableModel {
	return NewQAbstractTableModel(parent)
}
func NewQAbstractTableModel(parent QObject_ITF /*777 QObject **/) *QAbstractTableModel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTableModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractTableModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractTableModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:390
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractTableModel(QObject *)

/*

 */
func (*QAbstractTableModel) NewForInheritp() *QAbstractTableModel {
	return NewQAbstractTableModelp()
}
func NewQAbstractTableModelp() *QAbstractTableModel {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTableModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractTableModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractTableModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:391
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractTableModel()

/*

 */
func DeleteQAbstractTableModel(this *QAbstractTableModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTableModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:393
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
Returns the index of the item in the model specified by the given row, column and parent index.

When reimplementing this function in a subclass, call createIndex() to generate model indexes that other components can use to refer to items in your model.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also createIndex().
*/
func (this *QAbstractTableModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractTableModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:393
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
Returns the index of the item in the model specified by the given row, column and parent index.

When reimplementing this function in a subclass, call createIndex() to generate model indexes that other components can use to refer to items in your model.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also createIndex().
*/
func (this *QAbstractTableModel) Indexp(row int, column int) *QModelIndex /*123*/ {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractTableModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:394
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &) const

/*
Returns the sibling at row and column for the item at index, or an invalid QModelIndex if there is no sibling at that location.

sibling() is just a convenience function that finds the item's parent, and uses it to retrieve the index of the child item in the specified row and column.

This method can optionally be overridden for implementation-specific optimization.

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also index(), QModelIndex::row(), and QModelIndex::column().
*/
func (this *QAbstractTableModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if idx != nil && idx.QModelIndex_PTR() != nil {
		convArg2 = idx.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractTableModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:395
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
func (this *QAbstractTableModel) DropMimeData(data QMimeData_ITF /*777 const QMimeData **/, action int, row int, column int, parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg4 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractTableModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:398
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &) const

/*
Returns the item flags for the given index.

The base class implementation returns a combination of flags that enables the item (ItemIsEnabled) and allows it to be selected (ItemIsSelectable).

Note: This function can be invoked via the meta-object system and from QML. See Q_INVOKABLE.

See also Qt::ItemFlags.
*/
func (this *QAbstractTableModel) Flags(index QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractTableModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

//  body block end

//  keep block begin

func init_unused_10289() {
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
