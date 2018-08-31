package qtcore

// /usr/include/qt/QtCore/qabstractproxymodel.h
// #include <qabstractproxymodel.h>
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

// void resetInternalData()
func (this *QAbstractProxyModel) InheritResetInternalData(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "resetInternalData", f)
}

/*

 */
type QAbstractProxyModel struct {
	*QAbstractItemModel
}
type QAbstractProxyModel_ITF interface {
	QAbstractItemModel_ITF
	QAbstractProxyModel_PTR() *QAbstractProxyModel
}

func (ptr *QAbstractProxyModel) QAbstractProxyModel_PTR() *QAbstractProxyModel { return ptr }

func (this *QAbstractProxyModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemModel.GetCthis()
	}
}
func (this *QAbstractProxyModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemModel = NewQAbstractItemModelFromPointer(cthis)
}
func NewQAbstractProxyModelFromPointer(cthis unsafe.Pointer) *QAbstractProxyModel {
	bcthis0 := NewQAbstractItemModelFromPointer(cthis)
	return &QAbstractProxyModel{bcthis0}
}
func (*QAbstractProxyModel) NewFromPointer(cthis unsafe.Pointer) *QAbstractProxyModel {
	return NewQAbstractProxyModelFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractProxyModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractProxyModel(QObject *)

/*
Constructs a proxy model with the given parent.
*/
func (*QAbstractProxyModel) NewForInherit(parent QObject_ITF /*777 QObject **/) *QAbstractProxyModel {
	return NewQAbstractProxyModel(parent)
}
func NewQAbstractProxyModel(parent QObject_ITF /*777 QObject **/) *QAbstractProxyModel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractProxyModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractProxyModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractProxyModel(QObject *)

/*
Constructs a proxy model with the given parent.
*/
func (*QAbstractProxyModel) NewForInherit__() *QAbstractProxyModel {
	return NewQAbstractProxyModel__()
}
func NewQAbstractProxyModel__() *QAbstractProxyModel {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAbstractProxyModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAbstractProxyModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractProxyModel()

/*

 */
func DeleteQAbstractProxyModel(this *QAbstractProxyModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSourceModel(QAbstractItemModel *)

/*
Sets the given sourceModel to be processed by the proxy model.

Subclasses should call beginResetModel() at the beginning of the method, disconnect from the old model, call this method, connect to the new model, and call endResetModel().

Note: Setter function for property sourceModel.

See also sourceModel().
*/
func (this *QAbstractProxyModel) SetSourceModel(sourceModel QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 unsafe.Pointer
	if sourceModel != nil && sourceModel.QAbstractItemModel_PTR() != nil {
		convArg0 = sourceModel.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel14setSourceModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:63
// index:0
// Public Visibility=Default Availability=Available
// [8] QAbstractItemModel * sourceModel() const

/*
Returns the model that contains the data that is available through the proxy model.

Note: Getter function for property sourceModel.

See also setSourceModel().
*/
func (this *QAbstractProxyModel) SourceModel() *QAbstractItemModel /*777 QAbstractItemModel **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11sourceModelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractItemModelFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex mapToSource(const QModelIndex &) const

/*
Reimplement this function to return the model index in the source model that corresponds to the proxyIndex in the proxy model.

See also mapFromSource().
*/
func (this *QAbstractProxyModel) MapToSource(proxyIndex QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if proxyIndex != nil && proxyIndex.QModelIndex_PTR() != nil {
		convArg0 = proxyIndex.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11mapToSourceERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [24] QModelIndex mapFromSource(const QModelIndex &) const

/*
Reimplement this function to return the model index in the proxy model that corresponds to the sourceIndex from the source model.

See also mapToSource().
*/
func (this *QAbstractProxyModel) MapFromSource(sourceIndex QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if sourceIndex != nil && sourceIndex.QModelIndex_PTR() != nil {
		convArg0 = sourceIndex.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel13mapFromSourceERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QItemSelection mapSelectionToSource(const QItemSelection &) const

/*
Returns a source selection mapped from the specified proxySelection.

Reimplement this method to map proxy selections to source selections.
*/
func (this *QAbstractProxyModel) MapSelectionToSource(selection QItemSelection_ITF) *QItemSelection /*123*/ {
	var convArg0 unsafe.Pointer
	if selection != nil && selection.QItemSelection_PTR() != nil {
		convArg0 = selection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel20mapSelectionToSourceERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelection)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QItemSelection mapSelectionFromSource(const QItemSelection &) const

/*
Returns a proxy selection mapped from the specified sourceSelection.

Reimplement this method to map source selections to proxy selections.
*/
func (this *QAbstractProxyModel) MapSelectionFromSource(selection QItemSelection_ITF) *QItemSelection /*123*/ {
	var convArg0 unsafe.Pointer
	if selection != nil && selection.QItemSelection_PTR() != nil {
		convArg0 = selection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel22mapSelectionFromSourceERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelection)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool submit()

/*
Reimplemented from QAbstractItemModel::submit().
*/
func (this *QAbstractProxyModel) Submit() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel6submitEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void revert()

/*
Reimplemented from QAbstractItemModel::revert().
*/
func (this *QAbstractProxyModel) Revert() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel6revertEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const

/*
Reimplemented from QAbstractItemModel::data().

See also setData().
*/
func (this *QAbstractProxyModel) Data(proxyIndex QModelIndex_ITF, role int) *QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if proxyIndex != nil && proxyIndex.QModelIndex_PTR() != nil {
		convArg0 = proxyIndex.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const

/*
Reimplemented from QAbstractItemModel::data().

See also setData().
*/
func (this *QAbstractProxyModel) Data__(proxyIndex QModelIndex_ITF) *QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if proxyIndex != nil && proxyIndex.QModelIndex_PTR() != nil {
		convArg0 = proxyIndex.QModelIndex_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int) const

/*
Reimplemented from QAbstractItemModel::headerData().

See also setHeaderData().
*/
func (this *QAbstractProxyModel) HeaderData(section int, orientation int, role int) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:75
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int) const

/*
Reimplemented from QAbstractItemModel::headerData().

See also setHeaderData().
*/
func (this *QAbstractProxyModel) HeaderData__(section int, orientation int) *QVariant /*123*/ {
	// arg: 2, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::flags().
*/
func (this *QAbstractProxyModel) Flags(index QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)

/*
Reimplemented from QAbstractItemModel::setData().

See also data().
*/
func (this *QAbstractProxyModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)

/*
Reimplemented from QAbstractItemModel::setData().

See also data().
*/
func (this *QAbstractProxyModel) SetData__(index QModelIndex_ITF, value QVariant_ITF) bool {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setHeaderData(int, Qt::Orientation, const QVariant &, int)

/*
Reimplemented from QAbstractItemModel::setHeaderData().

See also headerData().
*/
func (this *QAbstractProxyModel) SetHeaderData(section int, orientation int, value QVariant_ITF, role int) bool {
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:81
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setHeaderData(int, Qt::Orientation, const QVariant &, int)

/*
Reimplemented from QAbstractItemModel::setHeaderData().

See also headerData().
*/
func (this *QAbstractProxyModel) SetHeaderData__(section int, orientation int, value QVariant_ITF) bool {
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::EditRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex buddy(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::buddy().
*/
func (this *QAbstractProxyModel) Buddy(index QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel5buddyERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canFetchMore(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::canFetchMore().
*/
func (this *QAbstractProxyModel) CanFetchMore(parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel12canFetchMoreERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void fetchMore(const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::fetchMore().
*/
func (this *QAbstractProxyModel) FetchMore(parent QModelIndex_ITF) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel9fetchMoreERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)

/*
Reimplemented from QAbstractItemModel::sort().
*/
func (this *QAbstractProxyModel) Sort(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)

/*
Reimplemented from QAbstractItemModel::sort().
*/
func (this *QAbstractProxyModel) Sort__(column int) {
	// arg: 1, Qt::SortOrder=Elaborated, Qt::SortOrder=Enum, , Invalid
	order := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:87
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize span(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::span().
*/
func (this *QAbstractProxyModel) Span(index QModelIndex_ITF) *QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel4spanERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSize)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::hasChildren().
*/
func (this *QAbstractProxyModel) HasChildren(parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::hasChildren().
*/
func (this *QAbstractProxyModel) HasChildren__() bool {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::sibling().
*/
func (this *QAbstractProxyModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if idx != nil && idx.QModelIndex_PTR() != nil {
		convArg2 = idx.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:92
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool canDropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::canDropMimeData().

This function was introduced in  Qt 5.4.
*/
func (this *QAbstractProxyModel) CanDropMimeData(data QMimeData_ITF /*777 const QMimeData **/, action int, row int, column int, parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg4 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel15canDropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::dropMimeData().

This function was introduced in  Qt 5.4.
*/
func (this *QAbstractProxyModel) DropMimeData(data QMimeData_ITF /*777 const QMimeData **/, action int, row int, column int, parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg4 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStringList mimeTypes() const

/*
Reimplemented from QAbstractItemModel::mimeTypes().
*/
func (this *QAbstractProxyModel) MimeTypes() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel9mimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDragActions() const

/*
Reimplemented from QAbstractItemModel::supportedDragActions().
*/
func (this *QAbstractProxyModel) SupportedDragActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel20supportedDragActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions() const

/*
Reimplemented from QAbstractItemModel::supportedDropActions().
*/
func (this *QAbstractProxyModel) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QAbstractProxyModel20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qabstractproxymodel.h:104
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void resetInternalData()

/*
Clears the roleNames of this proxy model.
*/
func (this *QAbstractProxyModel) ResetInternalData() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QAbstractProxyModel17resetInternalDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
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
