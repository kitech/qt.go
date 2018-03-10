package qtcore

// /usr/include/qt/QtCore/qidentityproxymodel.h
// #include <qidentityproxymodel.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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
type QIdentityProxyModel struct {
	*QAbstractProxyModel
}
type QIdentityProxyModel_ITF interface {
	QAbstractProxyModel_ITF
	QIdentityProxyModel_PTR() *QIdentityProxyModel
}

func (ptr *QIdentityProxyModel) QIdentityProxyModel_PTR() *QIdentityProxyModel { return ptr }

func (this *QIdentityProxyModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractProxyModel.GetCthis()
	}
}
func (this *QIdentityProxyModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractProxyModel = NewQAbstractProxyModelFromPointer(cthis)
}
func NewQIdentityProxyModelFromPointer(cthis unsafe.Pointer) *QIdentityProxyModel {
	bcthis0 := NewQAbstractProxyModelFromPointer(cthis)
	return &QIdentityProxyModel{bcthis0}
}
func (*QIdentityProxyModel) NewFromPointer(cthis unsafe.Pointer) *QIdentityProxyModel {
	return NewQIdentityProxyModelFromPointer(cthis)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QIdentityProxyModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIdentityProxyModel(QObject *)

/*
Constructs an identity model with the given parent.
*/
func NewQIdentityProxyModel(parent QObject_ITF /*777 QObject **/) *QIdentityProxyModel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIdentityProxyModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QIdentityProxyModel")
	return gothis
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIdentityProxyModel(QObject *)

/*
Constructs an identity model with the given parent.
*/
func NewQIdentityProxyModel__() *QIdentityProxyModel {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIdentityProxyModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QIdentityProxyModel")
	return gothis
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QIdentityProxyModel()

/*

 */
func DeleteQIdentityProxyModel(this *QIdentityProxyModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::columnCount().
*/
func (this *QIdentityProxyModel) ColumnCount(parent QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::columnCount().
*/
func (this *QIdentityProxyModel) ColumnCount__() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::index().
*/
func (this *QIdentityProxyModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::index().
*/
func (this *QIdentityProxyModel) Index__(row int, column int) *QModelIndex /*123*/ {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex mapFromSource(const QModelIndex &) const

/*
Reimplemented from QAbstractProxyModel::mapFromSource().
*/
func (this *QIdentityProxyModel) MapFromSource(sourceIndex QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if sourceIndex != nil && sourceIndex.QModelIndex_PTR() != nil {
		convArg0 = sourceIndex.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel13mapFromSourceERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex mapToSource(const QModelIndex &) const

/*
Reimplemented from QAbstractProxyModel::mapToSource().
*/
func (this *QIdentityProxyModel) MapToSource(proxyIndex QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if proxyIndex != nil && proxyIndex.QModelIndex_PTR() != nil {
		convArg0 = proxyIndex.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel11mapToSourceERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex parent(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::parent().
*/
func (this *QIdentityProxyModel) Parent(child QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if child != nil && child.QModelIndex_PTR() != nil {
		convArg0 = child.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel6parentERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::rowCount().
*/
func (this *QIdentityProxyModel) RowCount(parent QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::rowCount().
*/
func (this *QIdentityProxyModel) RowCount__() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int) const

/*
Reimplemented from QAbstractItemModel::headerData().
*/
func (this *QIdentityProxyModel) HeaderData(section int, orientation int, role int) *QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int) const

/*
Reimplemented from QAbstractItemModel::headerData().
*/
func (this *QIdentityProxyModel) HeaderData__(section int, orientation int) *QVariant /*123*/ {
	// arg: 2, int=Int, =Invalid,
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::dropMimeData().
*/
func (this *QIdentityProxyModel) DropMimeData(data QMimeData_ITF /*777 const QMimeData **/, action int, row int, column int, parent QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg4 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &) const

/*
Reimplemented from QAbstractItemModel::sibling().
*/
func (this *QIdentityProxyModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if idx != nil && idx.QModelIndex_PTR() != nil {
		convArg2 = idx.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QItemSelection mapSelectionFromSource(const QItemSelection &) const

/*
Reimplemented from QAbstractProxyModel::mapSelectionFromSource().
*/
func (this *QIdentityProxyModel) MapSelectionFromSource(selection QItemSelection_ITF) *QItemSelection /*123*/ {
	var convArg0 unsafe.Pointer
	if selection != nil && selection.QItemSelection_PTR() != nil {
		convArg0 = selection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel22mapSelectionFromSourceERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelection)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QItemSelection mapSelectionToSource(const QItemSelection &) const

/*
Reimplemented from QAbstractProxyModel::mapSelectionToSource().
*/
func (this *QIdentityProxyModel) MapSelectionToSource(selection QItemSelection_ITF) *QItemSelection /*123*/ {
	var convArg0 unsafe.Pointer
	if selection != nil && selection.QItemSelection_PTR() != nil {
		convArg0 = selection.QItemSelection_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel20mapSelectionToSourceERK14QItemSelection", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQItemSelectionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQItemSelection)
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QModelIndexList match(const QModelIndex &, int, const QVariant &, int, Qt::MatchFlags) const

/*
Reimplemented from QAbstractItemModel::match().
*/
func (this *QIdentityProxyModel) Match(start QModelIndex_ITF, role int, value QVariant_ITF, hits int, flags int) *QModelIndexList /*667*/ {
	var convArg0 unsafe.Pointer
	if start != nil && start.QModelIndex_PTR() != nil {
		convArg0 = start.QModelIndex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel5matchERK11QModelIndexiRK8QVarianti6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, convArg2, hits, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QModelIndexList match(const QModelIndex &, int, const QVariant &, int, Qt::MatchFlags) const

/*
Reimplemented from QAbstractItemModel::match().
*/
func (this *QIdentityProxyModel) Match__(start QModelIndex_ITF, role int, value QVariant_ITF) *QModelIndexList /*667*/ {
	var convArg0 unsafe.Pointer
	if start != nil && start.QModelIndex_PTR() != nil {
		convArg0 = start.QModelIndex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid,
	hits := int(1)
	// arg: 4, Qt::MatchFlags=Elaborated, Qt::MatchFlags=Typedef, QFlags<Qt::MatchFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel5matchERK11QModelIndexiRK8QVarianti6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, convArg2, hits, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QModelIndexList match(const QModelIndex &, int, const QVariant &, int, Qt::MatchFlags) const

/*
Reimplemented from QAbstractItemModel::match().
*/
func (this *QIdentityProxyModel) Match__1(start QModelIndex_ITF, role int, value QVariant_ITF, hits int) *QModelIndexList /*667*/ {
	var convArg0 unsafe.Pointer
	if start != nil && start.QModelIndex_PTR() != nil {
		convArg0 = start.QModelIndex_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	// arg: 4, Qt::MatchFlags=Elaborated, Qt::MatchFlags=Typedef, QFlags<Qt::MatchFlag>
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QIdentityProxyModel5matchERK11QModelIndexiRK8QVarianti6QFlagsIN2Qt9MatchFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role, convArg2, hits, flags)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQModelIndexListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setSourceModel(QAbstractItemModel *)

/*
Reimplemented from QAbstractProxyModel::setSourceModel().
*/
func (this *QIdentityProxyModel) SetSourceModel(sourceModel QAbstractItemModel_ITF /*777 QAbstractItemModel **/) {
	var convArg0 unsafe.Pointer
	if sourceModel != nil && sourceModel.QAbstractItemModel_PTR() != nil {
		convArg0 = sourceModel.QAbstractItemModel_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel14setSourceModelEP18QAbstractItemModel", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::insertColumns().
*/
func (this *QIdentityProxyModel) InsertColumns(column int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel13insertColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::insertColumns().
*/
func (this *QIdentityProxyModel) InsertColumns__(column int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel13insertColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::insertRows().
*/
func (this *QIdentityProxyModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::insertRows().
*/
func (this *QIdentityProxyModel) InsertRows__(row int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::removeColumns().
*/
func (this *QIdentityProxyModel) RemoveColumns(column int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel13removeColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::removeColumns().
*/
func (this *QIdentityProxyModel) RemoveColumns__(column int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel13removeColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::removeRows().
*/
func (this *QIdentityProxyModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qidentityproxymodel.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)

/*
Reimplemented from QAbstractItemModel::removeRows().
*/
func (this *QIdentityProxyModel) RemoveRows__(row int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QIdentityProxyModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
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
