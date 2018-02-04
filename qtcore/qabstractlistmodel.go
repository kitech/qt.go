package qtcore

// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"

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
		qtrt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin

type QAbstractListModel struct {
	*QAbstractItemModel
}

func (this *QAbstractListModel) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractItemModel.GetCthis()
	}
}
func (this *QAbstractListModel) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractItemModel = NewQAbstractItemModelFromPointer(cthis)
}
func NewQAbstractListModelFromPointer(cthis unsafe.Pointer) *QAbstractListModel {
	bcthis0 := NewQAbstractItemModelFromPointer(cthis)
	return &QAbstractListModel{bcthis0}
}
func (*QAbstractListModel) NewFromPointer(cthis unsafe.Pointer) *QAbstractListModel {
	return NewQAbstractListModelFromPointer(cthis)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:393
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QAbstractListModel) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractListModel10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:396
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAbstractListModel(QObject *)
func NewQAbstractListModel(parent *QObject /*777 QObject **/) *QAbstractListModel {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractListModelC1EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractListModelFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:397
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAbstractListModel()
func DeleteQAbstractListModel(this *QAbstractListModel) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractListModelD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:399
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex index(int, int, const QModelIndex &)
func (this *QAbstractListModel) Index(row int, column int, parent *QModelIndex) *QModelIndex /*123*/ {
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractListModel5indexEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:400
// index:0
// Public virtual Visibility=Default Availability=Available
// [24] QModelIndex sibling(int, int, const QModelIndex &)
func (this *QAbstractListModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex /*123*/ {
	var convArg2 = idx.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractListModel7siblingEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, column, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQModelIndexFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQModelIndex)
	return rv2
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:401
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool dropMimeData(const QMimeData *, Qt::DropAction, int, int, const QModelIndex &)
func (this *QAbstractListModel) DropMimeData(data *QMimeData /*777 const QMimeData **/, action int, row int, column int, parent *QModelIndex) bool {
	var convArg0 = data.GetCthis()
	var convArg4 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QAbstractListModel12dropMimeDataEPK9QMimeDataN2Qt10DropActionEiiRK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, action, row, column, convArg4)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:404
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::ItemFlags flags(const QModelIndex &)
func (this *QAbstractListModel) Flags(index *QModelIndex) int {
	var convArg0 = index.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QAbstractListModel5flagsERK11QModelIndex", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

//  body block end
