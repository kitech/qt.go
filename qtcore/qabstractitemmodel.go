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
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

// QModelIndex createIndex(int, int, void *)
func (this *QAbstractItemModel) InheritCreateIndex(f func(row int, column int, data unsafe.Pointer /*666*/) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "createIndex", f)
}

/*
 */
// size 16
type QAbstractItemModel struct {
	*QObject
}
type QAbstractItemModel_ITF interface {
	QObject_ITF
	QAbstractItemModel_PTR() *QAbstractItemModel
}

func (ptr *QAbstractItemModel) QAbstractItemModel_PTR() *QAbstractItemModel { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QAbstractItemModelFromptr(cthis Voidptr) *QAbstractItemModel {
	bcthis0 := QObjectFromptr(cthis)
	return &QAbstractItemModel{bcthis0}
}
func (*QAbstractItemModel) Fromptr(cthis Voidptr) *QAbstractItemModel {
	return QAbstractItemModelFromptr(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemModel(QObject *)

/*
 */
func (*QAbstractItemModel) NewForInherit(parent QObject_ITF /*777 QObject **/) *QAbstractItemModel {
	return NewQAbstractItemModel(parent)
}
func NewQAbstractItemModel(parent QObject_ITF /*777 QObject **/) *QAbstractItemModel {
	var convArg0 Voidptr
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(1538334395, "_ZN18QAbstractItemModelC2EP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QAbstractItemModelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractItemModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractItemModel(QObject *)

/*
 */
func (*QAbstractItemModel) NewForInheritp() *QAbstractItemModel {
	return NewQAbstractItemModelp()
}
func NewQAbstractItemModelp() *QAbstractItemModel {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 Voidptr
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc3(1538334395, "_ZN18QAbstractItemModelC2EP7QObject", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&cthis), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	gothis := QAbstractItemModelFromptr(cthis)
	qtrt.ConnectDestroyed(gothis, "QAbstractItemModel")
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:181
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool hasIndex(int, int, const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) HasIndex(row int, column int, parent QModelIndex_ITF) bool {
	var convArg2 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(281611678, "_ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&row), Voidptr(&column), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:181
// index:0
// Public Extend Visibility=Default Availability=Available
// [1] bool hasIndex(int, int, const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) HasIndexp(row int, column int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.Qtcc3(281611678, "_ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&row), Voidptr(&column), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:182
// index:0
// Public purevirtual virtual Indirect Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) Index(row int, column int, parent QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg2 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	sretobj := qtrt.Malloc(24) // QModelIndex
	rv, err := qtrt.Qtcc3(164512793, "_ZNK18QAbstractItemModel5indexEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr(), Voidptr(&row), Voidptr(&column), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := /*==*/ QModelIndexFromptr(rv.Ptr()) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:182
// index:0
// Public purevirtual virtual Indirect Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) Indexp(row int, column int) *QModelIndex /*123*/ {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	sretobj := qtrt.Malloc(24) // QModelIndex
	rv, err := qtrt.Qtcc3(164512793, "_ZNK18QAbstractItemModel5indexEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr(), Voidptr(&row), Voidptr(&column), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := /*==*/ QModelIndexFromptr(rv.Ptr()) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:184
// index:0
// Public purevirtual virtual Indirect Visibility=Default Availability=Available
// [24] QModelIndex parent(const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) Parent(child QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 Voidptr
	if child != nil && child.QModelIndex_PTR() != nil {
		convArg0 = child.QModelIndex_PTR().GetCthis()
	}
	sretobj := qtrt.Malloc(24) // QModelIndex
	rv, err := qtrt.Qtcc3(75286673, "_ZNK18QAbstractItemModel6parentERK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := /*==*/ QModelIndexFromptr(rv.Ptr()) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:186
// index:0
// Public virtual Indirect Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) Sibling(row int, column int, idx QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg2 Voidptr
	if idx != nil && idx.QModelIndex_PTR() != nil {
		convArg2 = idx.QModelIndex_PTR().GetCthis()
	}
	sretobj := qtrt.Malloc(24) // QModelIndex
	rv, err := qtrt.Qtcc3(783199811, "_ZNK18QAbstractItemModel7siblingEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr(), Voidptr(&row), Voidptr(&column), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := /*==*/ QModelIndexFromptr(rv.Ptr()) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:187
// index:0
// Public purevirtual virtual Direct Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) RowCount(parent QModelIndex_ITF) int {
	var convArg0 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1828919860, "_ZNK18QAbstractItemModel8rowCountERK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:187
// index:0
// Public purevirtual virtual Direct Visibility=Default Availability=Available
// [4] int rowCount(const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) RowCountp() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.Qtcc3(1828919860, "_ZNK18QAbstractItemModel8rowCountERK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:188
// index:0
// Public purevirtual virtual Direct Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) ColumnCount(parent QModelIndex_ITF) int {
	var convArg0 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(678361467, "_ZNK18QAbstractItemModel11columnCountERK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:188
// index:0
// Public purevirtual virtual Direct Visibility=Default Availability=Available
// [4] int columnCount(const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) ColumnCountp() int {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.Qtcc3(678361467, "_ZNK18QAbstractItemModel11columnCountERK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	return rv.Int() // 1111
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:189
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) HasChildren(parent QModelIndex_ITF) bool {
	var convArg0 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1032754129, "_ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:189
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool hasChildren(const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) HasChildrenp() bool {
	// arg: 0, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg0 = NewQModelIndex()
	rv, err := qtrt.Qtcc3(1032754129, "_ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:191
// index:0
// Public purevirtual virtual Indirect Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const

/*
 */
func (this *QAbstractItemModel) Data(index QModelIndex_ITF, role int) *QVariant /*123*/ {
	var convArg0 Voidptr
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	sretobj := qtrt.Malloc(16) // QVariant
	rv, err := qtrt.Qtcc3(2517699839, "_ZNK18QAbstractItemModel4dataERK11QModelIndexi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&sretobj), this.Addr(), Voidptr(&convArg0), Voidptr(&role))
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := /*==*/ QVariantFromptr(rv.Ptr()) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:191
// index:0
// Public purevirtual virtual Indirect Visibility=Default Availability=Available
// [16] QVariant data(const QModelIndex &, int) const

/*
 */
func (this *QAbstractItemModel) Datap(index QModelIndex_ITF) *QVariant /*123*/ {
	var convArg0 Voidptr
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	role := 0                  /*Qt::DisplayRole*/
	sretobj := qtrt.Malloc(16) // QVariant
	rv, err := qtrt.Qtcc3(2517699839, "_ZNK18QAbstractItemModel4dataERK11QModelIndexi", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&sretobj), this.Addr(), Voidptr(&convArg0), Voidptr(&role))
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := /*==*/ QVariantFromptr(rv.Ptr()) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:192
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)

/*
 */
func (this *QAbstractItemModel) SetData(index QModelIndex_ITF, value QVariant_ITF, role int) bool {
	var convArg0 Voidptr
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 Voidptr
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(279704600, "_ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&role))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:192
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool setData(const QModelIndex &, const QVariant &, int)

/*
 */
func (this *QAbstractItemModel) SetDatap(index QModelIndex_ITF, value QVariant_ITF) bool {
	var convArg0 Voidptr
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	var convArg1 Voidptr
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	role := 0 /*Qt::EditRole*/
	rv, err := qtrt.Qtcc3(279704600, "_ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&role))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:219
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) InsertRows(row int, count int, parent QModelIndex_ITF) bool {
	var convArg2 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(843062727, "_ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&row), Voidptr(&count), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:219
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool insertRows(int, int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) InsertRowsp(row int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.Qtcc3(843062727, "_ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&row), Voidptr(&count), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:220
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) InsertColumns(column int, count int, parent QModelIndex_ITF) bool {
	var convArg2 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(3123320753, "_ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&column), Voidptr(&count), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:220
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool insertColumns(int, int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) InsertColumnsp(column int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.Qtcc3(3123320753, "_ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&column), Voidptr(&count), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:221
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) RemoveRows(row int, count int, parent QModelIndex_ITF) bool {
	var convArg2 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(3635885257, "_ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&row), Voidptr(&count), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:221
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool removeRows(int, int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) RemoveRowsp(row int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.Qtcc3(3635885257, "_ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&row), Voidptr(&count), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:222
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) RemoveColumns(column int, count int, parent QModelIndex_ITF) bool {
	var convArg2 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg2 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(3609969157, "_ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&column), Voidptr(&count), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:222
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool removeColumns(int, int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) RemoveColumnsp(column int, count int) bool {
	// arg: 2, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg2 = NewQModelIndex()
	rv, err := qtrt.Qtcc3(3609969157, "_ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&column), Voidptr(&count), Voidptr(&convArg2))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:228
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool insertRow(int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) InsertRow(row int, parent QModelIndex_ITF) bool {
	var convArg1 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2197740031, "_ZN18QAbstractItemModel9insertRowEiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&row), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:228
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool insertRow(int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) InsertRowp(row int) bool {
	// arg: 1, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg1 = NewQModelIndex()
	rv, err := qtrt.Qtcc3(2197740031, "_ZN18QAbstractItemModel9insertRowEiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&row), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:229
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool insertColumn(int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) InsertColumn(column int, parent QModelIndex_ITF) bool {
	var convArg1 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2279385328, "_ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&column), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:229
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool insertColumn(int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) InsertColumnp(column int) bool {
	// arg: 1, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg1 = NewQModelIndex()
	rv, err := qtrt.Qtcc3(2279385328, "_ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&column), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:230
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool removeRow(int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) RemoveRow(row int, parent QModelIndex_ITF) bool {
	var convArg1 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(4007650377, "_ZN18QAbstractItemModel9removeRowEiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&row), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:230
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool removeRow(int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) RemoveRowp(row int) bool {
	// arg: 1, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg1 = NewQModelIndex()
	rv, err := qtrt.Qtcc3(4007650377, "_ZN18QAbstractItemModel9removeRowEiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&row), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:231
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool removeColumn(int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) RemoveColumn(column int, parent QModelIndex_ITF) bool {
	var convArg1 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg1 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1619948254, "_ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&column), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:231
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool removeColumn(int, const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) RemoveColumnp(column int) bool {
	// arg: 1, const QModelIndex &=LValueReference, QModelIndex=Record, , Invalid
	var convArg1 = NewQModelIndex()
	rv, err := qtrt.Qtcc3(1619948254, "_ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&column), Voidptr(&convArg1))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:232
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool moveRow(const QModelIndex &, int, const QModelIndex &, int)

/*
 */
func (this *QAbstractItemModel) MoveRow(sourceParent QModelIndex_ITF, sourceRow int, destinationParent QModelIndex_ITF, destinationChild int) bool {
	var convArg0 Voidptr
	if sourceParent != nil && sourceParent.QModelIndex_PTR() != nil {
		convArg0 = sourceParent.QModelIndex_PTR().GetCthis()
	}
	var convArg2 Voidptr
	if destinationParent != nil && destinationParent.QModelIndex_PTR() != nil {
		convArg2 = destinationParent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(3803393986, "_ZN18QAbstractItemModel7moveRowERK11QModelIndexiS2_i", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&convArg0), Voidptr(&sourceRow), Voidptr(&convArg2), Voidptr(&destinationChild))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:234
// index:0
// Public inline Extend Visibility=Default Availability=Available
// [1] bool moveColumn(const QModelIndex &, int, const QModelIndex &, int)

/*
 */
func (this *QAbstractItemModel) MoveColumn(sourceParent QModelIndex_ITF, sourceColumn int, destinationParent QModelIndex_ITF, destinationChild int) bool {
	var convArg0 Voidptr
	if sourceParent != nil && sourceParent.QModelIndex_PTR() != nil {
		convArg0 = sourceParent.QModelIndex_PTR().GetCthis()
	}
	var convArg2 Voidptr
	if destinationParent != nil && destinationParent.QModelIndex_PTR() != nil {
		convArg2 = destinationParent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(1894749172, "_ZN18QAbstractItemModel10moveColumnERK11QModelIndexiS2_i", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_POINTER, qtrt.FFITO_INT, this.Addr(), Voidptr(&convArg0), Voidptr(&sourceColumn), Voidptr(&convArg2), Voidptr(&destinationChild))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:237
// index:0
// Public virtual Ignore Visibility=Default Availability=Available
// [-2] void fetchMore(const QModelIndex &)

/*
 */
func (this *QAbstractItemModel) FetchMore(parent QModelIndex_ITF) {
	var convArg0 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(93984695, "_ZN18QAbstractItemModel9fetchMoreERK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:238
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool canFetchMore(const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) CanFetchMore(parent QModelIndex_ITF) bool {
	var convArg0 Voidptr
	if parent != nil && parent.QModelIndex_PTR() != nil {
		convArg0 = parent.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.Qtcc3(2951622094, "_ZNK18QAbstractItemModel12canFetchMoreERK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:241
// index:0
// Public virtual Indirect Visibility=Default Availability=Available
// [24] QModelIndex buddy(const QModelIndex &) const

/*
 */
func (this *QAbstractItemModel) Buddy(index QModelIndex_ITF) *QModelIndex /*123*/ {
	var convArg0 Voidptr
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	sretobj := qtrt.Malloc(24) // QModelIndex
	rv, err := qtrt.Qtcc3(1354301694, "_ZNK18QAbstractItemModel5buddyERK11QModelIndex", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr(), Voidptr(&convArg0))
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := /*==*/ QModelIndexFromptr(rv.Ptr()) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:299
// index:0
// Public virtual Extend Visibility=Default Availability=Available
// [1] bool submit()

/*
 */
func (this *QAbstractItemModel) Submit() bool {
	rv, err := qtrt.Qtcc3(2465187961, "_ZN18QAbstractItemModel6submitEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	return rv.Bool()
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:300
// index:0
// Public virtual Ignore Visibility=Default Availability=Available
// [-2] void revert()

/*
 */
func (this *QAbstractItemModel) Revert() {
	rv, err := qtrt.Qtcc3(380932091, "_ZN18QAbstractItemModel6revertEv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:311
// index:0
// Protected inline Indirect Visibility=Default Availability=Available
// [24] QModelIndex createIndex(int, int, void *) const

/*
 */
func (this *QAbstractItemModel) CreateIndex(row int, column int, data unsafe.Pointer /*666*/) *QModelIndex /*123*/ {
	sretobj := qtrt.Malloc(24) // QModelIndex
	rv, err := qtrt.Qtcc3(1298301286, "_ZNK18QAbstractItemModel11createIndexEiiPv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr(), Voidptr(&row), Voidptr(&column), Voidptr(&data))
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := /*==*/ QModelIndexFromptr(rv.Ptr()) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:311
// index:0
// Protected inline Indirect Visibility=Default Availability=Available
// [24] QModelIndex createIndex(int, int, void *) const

/*
 */
func (this *QAbstractItemModel) CreateIndexp(row int, column int) *QModelIndex /*123*/ {
	// arg: 2, void *=Pointer, =Invalid, , Invalid
	var data Voidptr
	sretobj := qtrt.Malloc(24) // QModelIndex
	rv, err := qtrt.Qtcc3(1298301286, "_ZNK18QAbstractItemModel11createIndexEiiPv", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr(), Voidptr(&row), Voidptr(&column), Voidptr(&data))
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := /*==*/ QModelIndexFromptr(rv.Ptr()) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:312
// index:1
// Protected inline Indirect Visibility=Default Availability=Available
// [24] QModelIndex createIndex(int, int, quintptr) const

/*
 */
func (this *QAbstractItemModel) CreateIndex1(row int, column int, id uint64) *QModelIndex /*123*/ {
	sretobj := qtrt.Malloc(24) // QModelIndex
	rv, err := qtrt.Qtcc3(2637747804, "_ZNK18QAbstractItemModel11createIndexEiiy", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, qtrt.FFITO_INT, qtrt.FFITO_POINTER, Voidptr(&sretobj), this.Addr(), Voidptr(&row), Voidptr(&column), Voidptr(&id))
	qtrt.ErrPrint3(err, rv)
	rv.High = uint64(uintptr(sretobj))
	rv2 := /*==*/ QModelIndexFromptr(rv.Ptr()) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

func DeleteQAbstractItemModel(this *QAbstractItemModel) {
	rv, err := qtrt.Qtcc3(23633593, "_ZN18QAbstractItemModelD2Ev", qtrt.FFITO_VOID, qtrt.FFITO_POINTER, this.Addr())
	qtrt.ErrPrint3(err, rv)
	//this.SetCthis(nil)
}

/*


 */
type QAbstractItemModel__LayoutChangeHint = int

//
const QAbstractItemModel__NoLayoutChangeHint QAbstractItemModel__LayoutChangeHint = 0

//
const QAbstractItemModel__VerticalSortHint QAbstractItemModel__LayoutChangeHint = 1

//
const QAbstractItemModel__HorizontalSortHint QAbstractItemModel__LayoutChangeHint = 2

func (this *QAbstractItemModel) LayoutChangeHintItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemModel_LayoutChangeHintItemName(val int) string {
	var nilthis *QAbstractItemModel
	return nilthis.LayoutChangeHintItemName(val)
}

/*


 */
type QAbstractItemModel__CheckIndexOption = int

//
const QAbstractItemModel__NoOption QAbstractItemModel__CheckIndexOption = 0

//
const QAbstractItemModel__IndexIsValid QAbstractItemModel__CheckIndexOption = 1

//
const QAbstractItemModel__DoNotUseParent QAbstractItemModel__CheckIndexOption = 2

//
const QAbstractItemModel__ParentIsInvalid QAbstractItemModel__CheckIndexOption = 4

func (this *QAbstractItemModel) CheckIndexOptionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QAbstractItemModel_CheckIndexOptionItemName(val int) string {
	var nilthis *QAbstractItemModel
	return nilthis.CheckIndexOptionItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10019() {
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
