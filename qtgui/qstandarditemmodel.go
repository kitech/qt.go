package qtgui

// /usr/include/qt/QtGui/qstandarditemmodel.h
// #include <qstandarditemmodel.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 83
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QStandardItemModel struct {
	*qtcore.QAbstractItemModel
}
type QStandardItemModel_ITF interface {
	qtcore.QAbstractItemModel_ITF
	QStandardItemModel_PTR() *QStandardItemModel
}

func (ptr *QStandardItemModel) QStandardItemModel_PTR() *QStandardItemModel { return ptr }

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
// [8] const QMetaObject * metaObject() const
func (this *QStandardItemModel) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:330
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStandardItemModel(QObject *)
func NewQStandardItemModel(parent qtcore.QObject_ITF /*777 QObject **/) *QStandardItemModel {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStandardItemModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStandardItemModel")
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:330
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QStandardItemModel(QObject *)
func NewQStandardItemModel__() *QStandardItemModel {
	// arg: 0, QObject *=Pointer, QObject=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModelC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStandardItemModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStandardItemModel")
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:331
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStandardItemModel(int, int, QObject *)
func NewQStandardItemModel_1(rows int, columns int, parent qtcore.QObject_ITF /*777 QObject **/) *QStandardItemModel {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModelC2EiiP7QObject", qtrt.FFI_TYPE_POINTER, rows, columns, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStandardItemModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStandardItemModel")
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:331
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QStandardItemModel(int, int, QObject *)
func NewQStandardItemModel_1_(rows int, columns int) *QStandardItemModel {
	// arg: 2, QObject *=Pointer, QObject=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModelC2EiiP7QObject", qtrt.FFI_TYPE_POINTER, rows, columns, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQStandardItemModelFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QStandardItemModel")
	return gothis
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:332
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QStandardItemModel()
func DeleteQStandardItemModel(this *QStandardItemModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:336
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const
func (this *QStandardItemModel) Index(row int, column int, parent qtcore.QModelIndex_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:336
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const
func (this *QStandardItemModel) Index__(row int, column int) *qtcore.QModelIndex /*123*/ {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:337
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex parent(const QModelIndex &) const
func (this *QStandardItemModel) Parent(child qtcore.QModelIndex_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if child != nil && child.QModelIndex_PTR() != nil {
		convArg0 = child.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel6parentERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:339
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const
func (this *QStandardItemModel) RowCount(parent qtcore.QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:339
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const
func (this *QStandardItemModel) RowCount__() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg0 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel8rowCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:340
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &) const
func (this *QStandardItemModel) ColumnCount(parent qtcore.QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:340
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &) const
func (this *QStandardItemModel) ColumnCount__() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg0 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel11columnCountERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:341
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &) const
func (this *QStandardItemModel) HasChildren(parent qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:341
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &) const
func (this *QStandardItemModel) HasChildren__() bool {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg0 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel11hasChildrenERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:343
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &) const
func (this *QStandardItemModel) Sibling(row int, column int, idx qtcore.QModelIndex_ITF) *qtcore.QModelIndex /*123*/ {
	var convArg2 unsafe.Pointer
	if idx != nil && idx.QModelIndex_PTR() != nil {
		convArg2 = idx.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:345
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const
func (this *QStandardItemModel) Data(index qtcore.QModelIndex_ITF, role int) *qtcore.QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:345
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const
func (this *QStandardItemModel) Data__(index qtcore.QModelIndex_ITF) *qtcore.QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid,
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel4dataERK11QModelIndexi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:346
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)
func (this *QStandardItemModel) SetData(index qtcore.QModelIndex_ITF, value qtcore.QVariant_ITF, role int) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:346
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)
func (this *QStandardItemModel) SetData__(index qtcore.QModelIndex_ITF, value qtcore.QVariant_ITF) bool {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid,
	role := 0 /*Qt::EditRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:348
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int) const
func (this *QStandardItemModel) HeaderData(section int, orientation int, role int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:348
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant headerData(int, Qt::Orientation, int) const
func (this *QStandardItemModel) HeaderData__(section int, orientation int) *qtcore.QVariant /*123*/ {
	// arg: 2, int=Int, =Invalid,
	role := 0 /*Qt::DisplayRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel10headerDataEiN2Qt11OrientationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, role)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:350
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setHeaderData(int, Qt::Orientation, const QVariant &, int)
func (this *QStandardItemModel) SetHeaderData(section int, orientation int, value qtcore.QVariant_ITF, role int) bool {
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:350
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool setHeaderData(int, Qt::Orientation, const QVariant &, int)
func (this *QStandardItemModel) SetHeaderData__(section int, orientation int, value qtcore.QVariant_ITF) bool {
	var convArg2 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg2 = value.QVariant_PTR().GetCthis()
	}
	// arg: 3, int=Int, =Invalid,
	role := 0 /*Qt::EditRole*/
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel13setHeaderDataEiN2Qt11OrientationERK8QVarianti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), section, orientation, convArg2, role)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:353
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)
func (this *QStandardItemModel) InsertRows(row int, count int, parent qtcore.QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:353
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)
func (this *QStandardItemModel) InsertRows__(row int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel10insertRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:354
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)
func (this *QStandardItemModel) InsertColumns(column int, count int, parent qtcore.QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel13insertColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:354
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)
func (this *QStandardItemModel) InsertColumns__(column int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel13insertColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:355
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)
func (this *QStandardItemModel) RemoveRows(row int, count int, parent qtcore.QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:355
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)
func (this *QStandardItemModel) RemoveRows__(row int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel10removeRowsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:356
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)
func (this *QStandardItemModel) RemoveColumns(column int, count int, parent qtcore.QModelIndex_ITF) bool {
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel13removeColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:356
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)
func (this *QStandardItemModel) RemoveColumns__(column int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg2 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel13removeColumnsEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, count, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:358
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &) const
func (this *QStandardItemModel) Flags(index qtcore.QModelIndex_ITF) int {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:359
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::DropActions supportedDropActions() const
func (this *QStandardItemModel) SupportedDropActions() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel20supportedDropActionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:364
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QStandardItemModel) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:368
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)
func (this *QStandardItemModel) Sort(column int, order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:368
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void sort(int, Qt::SortOrder)
func (this *QStandardItemModel) Sort__(column int) {
	// arg: 1, Qt::SortOrder=Elaborated, Qt::SortOrder=Enum,
	order := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel4sortEiN2Qt9SortOrderE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:370
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * itemFromIndex(const QModelIndex &) const
func (this *QStandardItemModel) ItemFromIndex(index qtcore.QModelIndex_ITF) *QStandardItem /*777 QStandardItem **/ {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel13itemFromIndexERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:371
// index:0
// Public Visibility=Default Availability=Available
// [24] QModelIndex indexFromItem(const QStandardItem *) const
func (this *QStandardItemModel) IndexFromItem(item QStandardItem_ITF /*777 const QStandardItem **/) *qtcore.QModelIndex /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg0 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel13indexFromItemEPK13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:373
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * item(int, int) const
func (this *QStandardItemModel) Item(row int, column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel4itemEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:373
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * item(int, int) const
func (this *QStandardItemModel) Item__(row int) *QStandardItem /*777 QStandardItem **/ {
	// arg: 1, int=Int, =Invalid,
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel4itemEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:374
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItem(int, int, QStandardItem *)
func (this *QStandardItemModel) SetItem(row int, column int, item QStandardItem_ITF /*777 QStandardItem **/) {
	var convArg2 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg2 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel7setItemEiiP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:375
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setItem(int, QStandardItem *)
func (this *QStandardItemModel) SetItem_1(row int, item QStandardItem_ITF /*777 QStandardItem **/) {
	var convArg1 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg1 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel7setItemEiP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:376
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * invisibleRootItem() const
func (this *QStandardItemModel) InvisibleRootItem() *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel17invisibleRootItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:378
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * horizontalHeaderItem(int) const
func (this *QStandardItemModel) HorizontalHeaderItem(column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel20horizontalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:379
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalHeaderItem(int, QStandardItem *)
func (this *QStandardItemModel) SetHorizontalHeaderItem(column int, item QStandardItem_ITF /*777 QStandardItem **/) {
	var convArg1 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg1 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel23setHorizontalHeaderItemEiP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:380
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * verticalHeaderItem(int) const
func (this *QStandardItemModel) VerticalHeaderItem(row int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel18verticalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:381
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalHeaderItem(int, QStandardItem *)
func (this *QStandardItemModel) SetVerticalHeaderItem(row int, item QStandardItem_ITF /*777 QStandardItem **/) {
	var convArg1 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg1 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel21setVerticalHeaderItemEiP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:383
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHorizontalHeaderLabels(const QStringList &)
func (this *QStandardItemModel) SetHorizontalHeaderLabels(labels qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if labels != nil && labels.QStringList_PTR() != nil {
		convArg0 = labels.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel25setHorizontalHeaderLabelsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:384
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVerticalHeaderLabels(const QStringList &)
func (this *QStandardItemModel) SetVerticalHeaderLabels(labels qtcore.QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if labels != nil && labels.QStringList_PTR() != nil {
		convArg0 = labels.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel23setVerticalHeaderLabelsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:386
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRowCount(int)
func (this *QStandardItemModel) SetRowCount(rows int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel11setRowCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rows)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:387
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColumnCount(int)
func (this *QStandardItemModel) SetColumnCount(columns int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel14setColumnCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:391
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void appendRow(QStandardItem *)
func (this *QStandardItemModel) AppendRow(item QStandardItem_ITF /*777 QStandardItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg0 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel9appendRowEP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:395
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void insertRow(int, QStandardItem *)
func (this *QStandardItemModel) InsertRow(row int, item QStandardItem_ITF /*777 QStandardItem **/) {
	var convArg1 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg1 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel9insertRowEiP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:397
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool insertRow(int, const QModelIndex &)
func (this *QStandardItemModel) InsertRow_1(row int, parent qtcore.QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel9insertRowEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:397
// index:1
// Public inline Visibility=Default Availability=Available
// [1] bool insertRow(int, const QModelIndex &)
func (this *QStandardItemModel) InsertRow_1_(row int) bool {
	// arg: 1, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg1 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel9insertRowEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:398
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool insertColumn(int, const QModelIndex &)
func (this *QStandardItemModel) InsertColumn(column int, parent qtcore.QModelIndex_ITF) bool {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel12insertColumnEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:398
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool insertColumn(int, const QModelIndex &)
func (this *QStandardItemModel) InsertColumn__(column int) bool {
	// arg: 1, const QModelIndex &=LValueReference, QModelIndex=Record,
	var convArg1 = qtcore.NewQModelIndex()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel12insertColumnEiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:400
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * takeItem(int, int)
func (this *QStandardItemModel) TakeItem(row int, column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel8takeItemEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:400
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * takeItem(int, int)
func (this *QStandardItemModel) TakeItem__(row int) *QStandardItem /*777 QStandardItem **/ {
	// arg: 1, int=Int, =Invalid,
	column := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel8takeItemEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:404
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * takeHorizontalHeaderItem(int)
func (this *QStandardItemModel) TakeHorizontalHeaderItem(column int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel24takeHorizontalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), column)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:405
// index:0
// Public Visibility=Default Availability=Available
// [8] QStandardItem * takeVerticalHeaderItem(int)
func (this *QStandardItemModel) TakeVerticalHeaderItem(row int) *QStandardItem /*777 QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel22takeVerticalHeaderItemEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:407
// index:0
// Public Visibility=Default Availability=Available
// [8] const QStandardItem * itemPrototype() const
func (this *QStandardItemModel) ItemPrototype() *QStandardItem /*777 const QStandardItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel13itemPrototypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQStandardItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:408
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemPrototype(const QStandardItem *)
func (this *QStandardItemModel) SetItemPrototype(item QStandardItem_ITF /*777 const QStandardItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg0 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel16setItemPrototypeEPK13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:414
// index:0
// Public Visibility=Default Availability=Available
// [4] int sortRole() const
func (this *QStandardItemModel) SortRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel8sortRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:415
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSortRole(int)
func (this *QStandardItemModel) SetSortRole(role int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel11setSortRoleEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), role)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:417
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QStringList mimeTypes() const
func (this *QStandardItemModel) MimeTypes() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QStandardItemModel9mimeTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:419
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QStandardItemModel) DropMimeData(data qtcore.QMimeData_ITF /*777 const QMimeData **/, action int, row int, column int, parent qtcore.QModelIndex_ITF) bool {
	var convArg0 unsafe.Pointer
	if data != nil && data.QMimeData_PTR() != nil {
		convArg0 = data.QMimeData_PTR().GetCthis()
	}
	var convArg4 unsafe.Pointer
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg4 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstandarditemmodel.h:422
// index:0
// Public Visibility=Default Availability=Available
// [-2] void itemChanged(QStandardItem *)
func (this *QStandardItemModel) ItemChanged(item QStandardItem_ITF /*777 QStandardItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QStandardItem_PTR() != nil {
		convArg0 = item.QStandardItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QStandardItemModel11itemChangedEP13QStandardItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
